package billing

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.com/bloom42/bloom/cmd/bloom/server/db"
	"gitlab.com/bloom42/lily/rz"
)

func FindInvoicesByGroupId(ctx context.Context, tx *sqlx.Tx, groupId string) ([]Invoice, error) {
	ret := []Invoice{}
	var err error
	logger := rz.FromCtx(ctx)

	query := `SELECT billing_invoices.* FROM billing_invoices
		INNER JOIN billing_customers ON billing_invoices.customer_id = billing_customers.id
		WHERE billing_customers.group_id = $1 ORDER BY created_at DESC`
	if tx == nil {
		err = db.DB.Select(&ret, query, groupId)
	} else {
		err = tx.Select(&ret, query, groupId)
	}
	if err != nil {
		logger.Error("finding invoices", rz.Err(err),
			rz.String("group.id", groupId))
		return ret, NewError(ErrorInvoiceNotFound)
	}

	return ret, nil
}
