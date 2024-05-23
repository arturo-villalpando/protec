package helpers

import (
	mo "protec/internal/models/commons"

	"math"
)

func Pagination(total int64, page int64, page_size int64, items any) mo.PaginationResponse {
	// Operations
	page_total := int64(math.Ceil(float64(total) / float64(page_size)))
	start := ((page - 1) * page_size) + 1
	end := page * page_size
	if page_total == page {
		end = total
	}
	// Response
	return mo.PaginationResponse{
		Total:     total,
		Page:      page,
		Size:      page_size,
		PageTotal: page_total,
		Start:     start,
		End:       end,
		Items:     items,
	}
}
