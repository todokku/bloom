package billing

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"gitlab.com/bloom42/lily/rz"
	"gitlab.com/bloom42/lily/uuid"
)

type PaymentMethod struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	IsDefault           bool   `json:"is_default" db:"is_default"`
	StripeID            string `json:"stripe_id" db:"stripe_id"`
	CardLast4           string `json:"card_last_4" db:"card_last_4"`
	CardExpirationMonth int64  `json:"card_expiration_month" db:"card_expiration_month"`
	CardExpirationYear  int64  `json:"card_expiration_year" db:"card_expiration_year"`

	CustomerID uuid.UUID `json:"customer_id" db:"customer_id"`
}

func FindPaymentMethodById(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (*PaymentMethod, error) {
	var ret *PaymentMethod
	var paymentMethod PaymentMethod
	var err error
	logger := rz.FromCtx(ctx)

	queryFind := "SELECT * FROM billing_payment_methods WHERE id = $1"
	err = tx.Get(&paymentMethod, queryFind, id)
	if err != nil {
		logger.Error("billing.FindPaymentMethodById: finding payment method", rz.Err(err),
			rz.String("payment_method.id", id.String()))
		return ret, NewError(ErrorPaymentMethodNotFound)
	}

	ret = &paymentMethod
	return ret, nil
}

func FindPaymentMethodByCustomer(ctx context.Context, tx *sqlx.Tx, customer *Customer, isDefault bool) (*PaymentMethod, error) {
	var ret *PaymentMethod
	var paymentMethod PaymentMethod
	var err error
	logger := rz.FromCtx(ctx)

	queryFind := "SELECT * FROM billing_payment_methods WHERE customer_id = $1 AND is_default = $2"
	err = tx.Get(&paymentMethod, queryFind, customer.ID, isDefault)
	if err != nil {
		logger.Error("billing.FindPaymentMethodByCustomer: finding payment method", rz.Err(err),
			rz.String("customer.id", customer.ID.String()))
		return ret, NewError(ErrorPaymentMethodNotFound)
	}

	ret = &paymentMethod
	return ret, nil
}
