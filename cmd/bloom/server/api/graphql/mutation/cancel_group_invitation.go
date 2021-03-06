package mutation

import (
	"context"

	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/apiutil"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/gqlerrors"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/model"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/db"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/domain/groups"
	"gitlab.com/bloom42/lily/rz"
)

func (r *Resolver) CancelGroupInvitation(ctx context.Context, input model.CancelGroupInvitationInput) (bool, error) {
	ret := false
	logger := rz.FromCtx(ctx)
	currentUser := apiutil.UserFromCtx(ctx)

	if currentUser == nil {
		return ret, gqlerrors.AuthenticationRequired()
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		logger.Error("mutation.CancelGroupInvitation: Starting transaction", rz.Err(err))
		return ret, gqlerrors.New(groups.NewError(groups.ErrorCancelingInvitation))
	}

	var invitation groups.Invitation

	queryGetInvitation := "SELECT * FROM groups_invitations WHERE id = $1 FOR UPDATE"
	err = tx.Get(&invitation, queryGetInvitation, input.ID)
	if err != nil {
		tx.Rollback()
		logger.Error("mutation.CancelGroupInvitation: fetching invitation", rz.Err(err),
			rz.String("invitation.id", input.ID.String()))
		return ret, gqlerrors.New(groups.NewError(groups.ErrorInvitationNotFound))
	}

	err = groups.CancelInvitation(ctx, tx, *currentUser, invitation)
	if err != nil {
		tx.Rollback()
		return ret, gqlerrors.New(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("mutation.CancelGroupInvitation: Committing transaction", rz.Err(err))
		return ret, gqlerrors.New(groups.NewError(groups.ErrorCancelingInvitation))
	}

	return ret, nil
}
