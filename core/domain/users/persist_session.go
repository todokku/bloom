package users

import (
	"context"
	"encoding/json"

	"gitlab.com/bloom42/bloom/core/api/model"
	"gitlab.com/bloom42/bloom/core/db"
	"gitlab.com/bloom42/bloom/core/domain/preferences"
)

func PersistSession(signin *model.SignedIn) error {
	meData, err := json.Marshal(signin.Me)
	if err != nil {
		return err
	}

	sessionData, err := json.Marshal(signin.Session)
	if err != nil {
		return err
	}

	tx, err := db.DB.Beginx()
	if err != nil {
		return err
	}

	err = preferences.Set(context.Background(), tx, "me", string(meData))
	if err != nil {
		tx.Rollback()
		return err
	}

	err = preferences.Set(context.Background(), tx, "session", string(sessionData))
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
