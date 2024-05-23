package reports

import (
	"fmt"
	"net/http"
	"strconv"

	"protec/internal/db"
	mc "protec/internal/models/commons"
	mo "protec/internal/models/reports"

	"github.com/gin-gonic/gin"
)

// Read godoc
// @Summary Read
// @Description Read
// @Tags Reports Transactions
// @Accept json
// @Produce json
// @Success      200  {object}  mo.Earns
// @Failure      400  {object}  mc.HTTPError
// @Failure      404  {object}  mc.HTTPError
// @Failure      500  {object}  mc.HTTPError
// @Security Bearer
// @securityDefinitions.apikey Bearer
// @in header
// @name authorization
// @description	Type "Bearer" followed by a space and JWT token.
// @Router /report/transactions/ [get]
func (s *ReportController) TransactionsEarns(ctx *gin.Context) {
	var _ mc.HTTPSuccess
	// Database connection and query
	q := db.New(s.dbConn)
	// Read commission from merchant
	merchant, err := q.TransactionsEarns(ctx)
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error reading data",
		})
		return
	}
	// Response
	ctx.JSON(http.StatusCreated, mo.Earns{Earns: merchant})
}

// Read Transaction Merchants Earns godoc
// @Summary Read Transaction Merchants Earns
// @Description Read Transaction Merchants Earns
// @Tags Reports Transactions
// @Param   id     path    int32     true        "1"
// @Accept json
// @Produce json
// @Success      200  {object}  mo.Earns
// @Failure      400  {object}  mc.HTTPError
// @Failure      404  {object}  mc.HTTPError
// @Failure      500  {object}  mc.HTTPError
// @Security Bearer
// @securityDefinitions.apikey Bearer
// @in header
// @name authorization
// @description	Type "Bearer" followed by a space and JWT token.
// @Router /report/transactions/{id} [get]
func (s *ReportController) TransactionsMerchantsEarns(ctx *gin.Context) {
	var _ mc.HTTPSuccess
	// Validation
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusBadRequest))
		// Response
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	// Database connection and query
	q := db.New(s.dbConn)
	// Read commission from merchant
	merchant, err := q.TransactionsEarnsByMerchant(ctx, int32(id))
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error reading data",
		})
		return
	}
	// Response
	ctx.JSON(http.StatusCreated, mo.Earns{Earns: merchant})
}
