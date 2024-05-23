package transactions

import (
	"github.com/gin-gonic/gin"
)

func (r *TransactionRouter) TransactionRouter(router *gin.Engine) {
	router.POST("/api/transaction/", r.controller.CreateTransaction)
}
