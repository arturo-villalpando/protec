package reports

import (
	"github.com/gin-gonic/gin"
)

func (r *ReportRouter) ReportTransactionRouter(router *gin.Engine) {
	router.GET("/api/report/transactions/", r.controller.TransactionsEarns)
	router.GET("/api/report/transactions/:id", r.controller.TransactionsMerchantsEarns)
}
