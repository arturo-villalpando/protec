package merchants

import (
	"protec/internal/db"
	mo "protec/internal/models/merchants"
)

func getMerchant(req db.Merchant) mo.Merchant {
	return mo.Merchant{
		MerchantId:   req.MerchantID,
		MerchantName: req.MerchantName,
		Commission:   req.Commission,
		CreatedAt:    req.CreatedAt.Time, // req.CreatedAt.Time.Format("0102061504")
		UpdatedAt:    req.UpdatedAt.Time, // req.CreatedAt.Time.Format("0102061504")
	}
}
