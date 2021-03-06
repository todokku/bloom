package admin

import (
	"context"

	"gitlab.com/bloom42/bloom/core/api"
	"gitlab.com/bloom42/lily/graphql"
)

func FetchDashboardData() (DashboardData, error) {
	client := api.Client()
	var ret DashboardData

	req := graphql.NewRequest(`
	query {
		metadata {
			os
			arch
			version
			gitCommit
		}
		users {
			totalCount
		}
	}
	`)

	err := client.Do(context.Background(), req, &ret)

	return ret, err
}
