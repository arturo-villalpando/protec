package helpers

import (
	mo "protec/internal/models/commons"

	"github.com/gin-gonic/gin"
)

func NewError(ctx *gin.Context, status int, err error) {
	er := mo.HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}
