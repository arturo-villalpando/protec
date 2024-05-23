package models

import "time"

type Merchant struct {
	MerchantId   int32     `json:"merchant_id" form:"merchant_id" binding:"required,number,gt=0" example:"1"`
	MerchantName string    `json:"merchant_name" form:"merchant_name" binding:"required,max=128" example:"Qwerty"`
	Commission   float64   `json:"commission" form:"commission" binding:"required,numeric,gt=0,lte=100" example:"1"`
	CreatedAt    time.Time `bson:"created_at" binding:"required,datetime" example:"2022-12-10T12:00:00Z"`
	UpdatedAt    time.Time `bson:"updated_at" binding:"required,datetime" example:"2022-12-10T12:00:00Z"`
}

type MerchantCreate struct {
	MerchantName string  `json:"merchant_name" form:"merchant_name" binding:"required,max=128" example:"Qwerty"`
	Commission   float64 `json:"commission" form:"commission" binding:"required,numeric,gt=0,lte=100" example:"1"`
	CreatedAt    string  `json:"created_at" form:"created_at" binding:"required" example:"2021-01-01 00:00:00"`
	UpdatedAt    string  `json:"updated_at" form:"updated_at" binding:"required" example:"2021-01-01 00:00:00"`
}

type MerchantRequest struct {
	MerchantName string  `json:"merchant_name" form:"merchant_name" binding:"required,max=128" example:"Qwerty"`
	Commission   float64 `json:"commission" form:"commission" binding:"required,numeric,gt=0,lte=100" example:"1"`
}
