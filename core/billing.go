package core

import (
	"C"
	"encoding/json"

	"gitlab.com/bloom42/bloom/core/api/model"
	"gitlab.com/bloom42/bloom/core/domain/billing"
	"gitlab.com/bloom42/bloom/core/domain/kernel"
)

func handleBillingMehtod(method string, jsonParams json.RawMessage) MessageOut {
	switch method {
	case "fetchMyProfile":
		res, err := billing.FetchMyProfile()
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	case "fetchGroupProfile":
		var params billing.FetchGroupProfileParams
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		res, err := billing.FetchGroupProfile(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	case "fetchPlans":
		res, err := billing.FetchPlans()
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	case "deletePlan":
		var params model.DeleteBillingPlanInput
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		err = billing.DeletePlan(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: kernel.Empty{}}
	case "createPlan":
		var params model.BillingPlanInput
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		res, err := billing.CreatePlan(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	case "updatePlan":
		var params model.BillingPlanInput
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		res, err := billing.UpdatePlan(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	case "updateSubscription":
		var params model.UpdateBillingSubscriptionInput
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		res, err := billing.UpdateSubscription(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	case "addPaymentMethod":
		var params billing.AddPaymentMethodParams
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		res, err := billing.AddPaymentMethod(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	case "removePaymentMethod":
		var params model.RemovePaymentMethodInput
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		err = billing.RemovePaymentMethod(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: kernel.Empty{}}
	case "changeDefaultPaymentMethod":
		var params model.ChangeDefaultPaymentMethodInput
		err := json.Unmarshal(jsonParams, &params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		res, err := billing.ChangeDefaultPaymentMethod(params)
		if err != nil {
			return InternalError(err) // TODO(z0mbie42): return error
		}
		return MessageOut{Data: res}
	default:
		return methodNotFoundError(method, "billing")
	}
}
