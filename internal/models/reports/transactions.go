package models

type Earns struct {
	Earns float64 `form:"earns" json:"earns" xml:"earns" binding:"required,number,gt=0" example:"1"`
}
