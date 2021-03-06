package mutation

import (
	"context"

	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/apiutil"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/gqlerrors"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/api/graphql/model"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/domain/billing"
)

func (r *Resolver) UpdateBillingSubscription(ctx context.Context, input model.UpdateBillingSubscriptionInput) (*model.BillingSubscription, error) {
	var ret *model.BillingSubscription
	var err error
	currentUser := apiutil.UserFromCtx(ctx)

	if currentUser == nil {
		return ret, gqlerrors.AuthenticationRequired()
	}

	customer, newPlan, err := billing.ChangeSubscription(ctx, currentUser, input.UserID, input.GroupID, input.PlanID)
	if err != nil {
		return ret, gqlerrors.New(err)
	}

	ret = &model.BillingSubscription{
		Plan: &model.BillingPlan{
			ID:          newPlan.ID,
			Name:        newPlan.Name,
			Description: newPlan.Description,
			Product:     model.BillingProduct(newPlan.Product),
			Price:       newPlan.Price,
			IsPublic:    newPlan.IsPublic,
			Storage:     newPlan.Storage,
		},
		UpdatedAt:   customer.SubscriptionUpdatedAt,
		UsedStorage: customer.UsedStorage,
	}
	return ret, nil
}
