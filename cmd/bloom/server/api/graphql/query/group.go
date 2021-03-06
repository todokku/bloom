package query

import (
	"context"

	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/apiutil"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/gqlerrors"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/model"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/domain/groups"
	"gitlab.com/bloom42/lily/uuid"
)

// Group find a group
func (r *Resolver) Group(ctx context.Context, id uuid.UUID) (*model.Group, error) {
	var ret *model.Group
	currentUser := apiutil.UserFromCtx(ctx)

	if currentUser == nil {
		return ret, gqlerrors.AuthenticationRequired()
	}

	err := groups.CheckUserIsGroupMemberNoTx(ctx, currentUser.ID, id)
	if err != nil && !currentUser.IsAdmin {
		return ret, gqlerrors.New(err)
	}

	group, err := groups.FindGroupById(ctx, nil, id)
	if err != nil {
		return ret, gqlerrors.New(err)
	}

	ret = &model.Group{
		ID:          &group.ID,
		CreatedAt:   &group.CreatedAt,
		Name:        group.Name,
		Description: group.Description,
		AvatarURL:   nil,
	}

	return ret, nil
}
