package handler

import (
	"context"

	"github.com/twitchtv/twirp"
	rpc "gitlab.com/bloom42/bloom/common/rpc/groups"
	"gitlab.com/bloom42/bloom/server/api/apictx"
	"gitlab.com/bloom42/bloom/server/db"
	"gitlab.com/bloom42/bloom/server/domain/groups"
	"gitlab.com/bloom42/libs/rz-go"
)

func (handler Handler) CreateGroup(ctx context.Context, params *rpc.CreateGroupParams) (*rpc.Group, error) {
	ret := &rpc.Group{}

	logger := rz.FromCtx(ctx)
	apiCtx, ok := ctx.Value(apictx.Key).(*apictx.Context)
	if !ok {
		return ret, twirp.InternalError("internal error")
	}
	if apiCtx.AuthenticatedAccount == nil {
		twerr := twirp.NewError(twirp.Unauthenticated, "authentication required")
		return ret, twerr
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		logger.Error("groups.CreateGroup: Starting transaction", rz.Err(err))
		return ret, twirp.InternalError(groups.ErrorCreateGroupMsg)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("groups.CreateGroup: Committing transaction", rz.Err(err))
		return ret, twirp.InternalError(groups.ErrorCreateGroupMsg)
	}

	return ret, nil
}