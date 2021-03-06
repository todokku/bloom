package users

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/db"
	"gitlab.com/bloom42/lily/rz"
	"gitlab.com/bloom42/lily/uuid"
)

func FindUserById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (*User, error) {
	ret := &User{}
	var err error
	logger := rz.FromCtx(ctx)

	queryFind := "SELECT * FROM users WHERE id = $1"
	if tx == nil {
		err = db.DB.Get(ret, queryFind, id)
	} else {
		err = tx.Get(ret, queryFind, id)
	}
	if err != nil {
		logger.Error("finding user", rz.Err(err),
			rz.String("user.id", id.String()))
		return ret, NewError(ErrorUserNotFound)
	}

	return ret, err
}
