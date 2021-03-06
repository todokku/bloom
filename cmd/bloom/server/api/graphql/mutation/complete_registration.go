package mutation

import (
	"context"
	"time"

	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/apiutil"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/gqlerrors"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/model"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/db"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/domain/billing"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/domain/users"
	"gitlab.com/bloom42/lily/crypto"
	"gitlab.com/bloom42/lily/rz"
)

func (r *Resolver) CompleteRegistration(ctx context.Context, input model.CompleteRegistrationInput) (*model.SignedIn, error) {
	var ret *model.SignedIn
	logger := rz.FromCtx(ctx)
	currentUser := apiutil.UserFromCtx(ctx)
	apiCtx := apiutil.ApiCtxFromCtx(ctx)
	if apiCtx == nil {
		logger.Error("mutation.CompleteRegistration: error getting apiCtx from context")
		return ret, gqlerrors.Internal()
	}

	if currentUser != nil {
		return ret, gqlerrors.MustNotBeAuthenticated()
	}

	// sleep to prevent spam and bruteforce
	sleep, err := crypto.RandInt64(500, 800)
	if err != nil {
		logger.Error("mutation.CompleteRegistration: generating random int", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	tx, err := db.DB.Beginx()
	if err != nil {
		logger.Error("mutation.CompleteRegistration: Starting transaction", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}

	// find pending user
	var pendingUser users.PendingUser
	err = tx.Get(&pendingUser, "SELECT * FROM pending_users WHERE id = $1 FOR UPDATE", input.ID)
	if err != nil {
		tx.Rollback()
		logger.Error("mutation.CompleteRegistration: getting pending user", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}

	// delete pending user
	err = users.DeletePendingUser(ctx, tx, pendingUser.ID.String())
	if err != nil {
		tx.Rollback()
		return ret, gqlerrors.New(err)
	}

	// create user
	createUserParams := users.CreateUserParams{
		PendingUser:         pendingUser,
		Username:            input.Username,
		AuthKey:             input.AuthKey,
		PublicKey:           input.PublicKey,
		EncryptedPrivateKey: input.EncryptedPrivateKey,
		PrivateKeyNonce:     input.PrivateKeyNonce,
	}
	newUser, err := users.CreateUser(ctx, tx, createUserParams)
	if err != nil {
		tx.Rollback()
		return ret, gqlerrors.New(err)
	}

	// create customer profile
	_, err = billing.CreateCustomer(ctx, tx, &newUser, &newUser.ID, nil)
	if err != nil {
		tx.Rollback()
		return ret, gqlerrors.New(err)
	}

	// start session
	device := users.SessionDevice{
		OS:   input.Device.Os.String(),
		Type: input.Device.Type.String(),
	}

	newSession, token, err := users.StartSession(ctx, tx, newUser.ID, device)
	if err != nil {
		tx.Rollback()
		return ret, gqlerrors.New(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("users.VerifyRegistration: Committing transaction", rz.Err(err))
		return ret, gqlerrors.New(users.NewError(users.ErrorCreatingPendingUser))
	}

	ret = &model.SignedIn{
		Session: &model.Session{
			ID:    newSession.ID,
			Token: &token,
			Device: &model.SessionDevice{
				Os:   model.SessionDeviceOs(device.OS),
				Type: model.SessionDeviceType(device.Type),
			},
		},
		Me: &model.User{
			ID:          &newUser.ID,
			AvatarURL:   nil,
			CreatedAt:   &newUser.CreatedAt,
			Username:    newUser.Username,
			FirstName:   &newUser.FirstName,
			LastName:    &newUser.LastName,
			DisplayName: newUser.DisplayName,
			IsAdmin:     newUser.IsAdmin,
		},
	}
	return ret, nil
}
