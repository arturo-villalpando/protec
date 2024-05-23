package transactions

import (
	"protec/internal/db"
	mo "protec/internal/models/transactions"
)

func getTransaction(req db.Transaction) mo.Transaction {
	return mo.Transaction{
		TransactionId: req.TransactionID,
		MerchantId:    req.MerchantID,
		Amount:        req.Amount,
		Commission:    req.Commission,
		Fee:           req.Fee,
		CreatedAt:     req.CreatedAt.Time, // req.CreatedAt.Time.Format("0102061504")
	}
}
