package models

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Created successfully"`
}

type SearchRequest struct {
	Search string `form:"search" json:"search" xml:"search" binding:"required,max=128" example:"musica"`
}

type IntRequest struct {
	Id int32 `form:"id" json:"id" xml:"id" binding:"required,number,gt=0" example:"1"`
}

type PaginationRequest struct {
	Page      int64  `form:"page" json:"page" xml:"page" binding:"required,number,gt=0" example:"1"`
	Limit     int64  `form:"limit,default=10" json:"limit" xml:"limit" binding:"required,number,gt=0" example:"10"`
	Search    string `form:"search" json:"search" xml:"search" binding:"omitempty" example:"nombre"`
	SortBy    string `form:"sort_by,default=created_at" json:"sort_by" xml:"sort_by" binding:"omitempty,oneof=name created_at" example:"created_at"`
	SortOrder string `form:"sort_order,default=desc" json:"sort_order" xml:"sort_order" binding:"omitempty,oneof=asc desc" example:"desc"`
}

type PaginationResponse struct {
	Total     int64 `form:"total,default=0" json:"total" xml:"total" binding:"required,number,gt=0" example:"0"`
	Page      int64 `form:"page,default=1" json:"page" xml:"page" binding:"required,number,gt=0" example:"1"`
	Size      int64 `form:"size,default=10" json:"size" xml:"size" binding:"required,number,gt=0" example:"10"`
	PageTotal int64 `form:"page_total,default=0" json:"page_total" xml:"page_total" binding:"required,number,gt=0" example:"0"`
	Start     int64 `form:"start,default=1" json:"start" xml:"start" binding:"required,number,gt=0" example:"1"`
	End       int64 `form:"end,default=1" json:"end" xml:"end" binding:"required,number,gt=0" example:"1"`
	Items     any   `form:"items" json:"items" xml:"items" binding:"required,max=128" example:"Pagos"`
}
