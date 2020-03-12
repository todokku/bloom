package groups

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.com/bloom42/bloom/server/db"
	"gitlab.com/bloom42/libs/rz-go"
)

func FindGroupMembers(ctx context.Context, tx *sqlx.Tx, groupId string) ([]Member, error) {
	ret := []Member{}
	var err error
	logger := rz.FromCtx(ctx)

	query := `SELECT DISTINCT users.id, users.*, groups_members.role, groups_members.joined_at FROM groups, users
		INNER JOIN groups_members ON groups_members.user_id = users.id
		WHERE groups_members.group_id = $1`
	if tx == nil {
		err = db.DB.Select(&ret, query, groupId)
	} else {
		err = tx.Select(&ret, query, groupId)
	}
	if err != nil {
		logger.Error("finding group members", rz.Err(err),
			rz.String("group.id", groupId))
		return ret, NewError(ErrorGroupNotFound)
	}

	return ret, err
}