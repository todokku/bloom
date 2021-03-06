package groups

import (
	"context"

	"gitlab.com/bloom42/bloom/core/api"
	"gitlab.com/bloom42/bloom/core/api/model"
	"gitlab.com/bloom42/bloom/core/db"
	"gitlab.com/bloom42/lily/graphql"
)

func QuitGroup(input model.QuitGroupInput) (bool, error) {
	client := api.Client()

	var resp struct {
		Success bool `json:"quitGroup"`
	}
	req := graphql.NewRequest(`
	mutation($input: QuitGroupInput!) {
		quitGroup(input: $input)
	}
	`)
	req.Var("input", input)

	err := client.Do(context.Background(), req, &resp)
	if err == nil {
		_, err = db.DB.Exec("DELETE FROM groups WHERE id = ?", input.ID)
	}

	return resp.Success, err
}
