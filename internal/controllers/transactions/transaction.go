package transactions

import (
	"fmt"
	"net/http"

	"protec/internal/db"
	mc "protec/internal/models/commons"
	mo "protec/internal/models/transactions"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create
// @Description Create
// @Tags Transactions
// @Param req body mo.TransactionRequest true "Create"
// @Accept json
// @Produce json
// @Success      200  {object}  mo.Transaction
// @Failure      400  {object}  mc.HTTPError
// @Failure      404  {object}  mc.HTTPError
// @Failure      500  {object}  mc.HTTPError
// @Security Bearer
// @securityDefinitions.apikey Bearer
// @in header
// @name authorization
// @description	Type "Bearer" followed by a space and JWT token.
// @Router /transaction/ [post]
func (s *TransactionController) CreateTransaction(ctx *gin.Context) {
	var _ mc.HTTPSuccess
	// Validation
	var req mo.TransactionRequest
	err := ctx.ShouldBindJSON(&req)
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
	merchant, err := q.ReadMerchant(ctx, req.MerchantId)
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error reading merchant",
		})
		return
	}
	// Create
	fee := req.Amount * (merchant.Commission / 100.0)
	item, err := q.CreateTransaction(ctx, db.CreateTransactionParams{
		MerchantID: req.MerchantId,
		Amount:     req.Amount,
		Commission: merchant.Commission,
		Fee:        fee,
	})
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error creating Transaction",
		})
		return
	}
	// Response
	ctx.JSON(http.StatusCreated, getTransaction(item))
}
