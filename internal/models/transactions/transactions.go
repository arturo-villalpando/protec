package models

import "time"

type Transaction struct {
	TransactionId int32     `json:"transaction_id" form:"transaction_id" binding:"required,number,gt=0" example:"1"`
	MerchantId    int32     `json:"merchant_id" form:"merchant_id" binding:"required,number,gt=0" example:"1"`
	Amount        float64   `json:"amount" form:"amount" binding:"required,numeric,gt=0" example:"100.00"`
	Commission    float64   `json:"commission" form:"commission" binding:"required,numeric,gt=0,lte=100" example:"1"`
	Fee           float64   `json:"fee" form:"fee" binding:"required,numeric,gt=0" example:"1.00"`
	CreatedAt     time.Time `bson:"created_at" binding:"required,datetime" example:"2022-12-10T12:00:00Z"`
}

type TransactionCreate struct {
	MerchantId int32   `json:"merchant_id" form:"merchant_id" binding:"required,number,gt=0" example:"1"`
	Amount     float64 `json:"amount" form:"amount" binding:"required,number,gt=0" example:"100.00"`
	Commission float64 `json:"commission" form:"commission" binding:"required,numeric,gt=0,lte=100" example:"1"`
	Fee        float64 `json:"fee" form:"fee" binding:"required,numeric,gt=0" example:"1.00"`
	CreatedAt  string  `json:"created_at" form:"created_at" binding:"required" example:"2021-01-01 00:00:00"`
}

type TransactionRequest struct {
	MerchantId int32   `json:"merchant_id" form:"merchant_id" binding:"required,number,gt=0" example:"1"`
	Amount     float64 `json:"amount" form:"amount" binding:"required,numeric,gt=0" example:"100.00"`
}
