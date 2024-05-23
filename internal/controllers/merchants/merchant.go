package merchants

import (
	"fmt"
	"net/http"
	"strconv"

	"protec/internal/db"
	mc "protec/internal/models/commons"
	mo "protec/internal/models/merchants"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create
// @Description Create
// @Tags Merchants
// @Param req body mo.MerchantRequest true "Create"
// @Accept json
// @Produce json
// @Success      200  {object}  mo.Merchant
// @Failure      400  {object}  mc.HTTPError
// @Failure      404  {object}  mc.HTTPError
// @Failure      500  {object}  mc.HTTPError
// @Security Bearer
// @securityDefinitions.apikey Bearer
// @in header
// @name authorization
// @description	Type "Bearer" followed by a space and JWT token.
// @Router /merchant/ [post]
func (s *MerchantController) CreateMerchant(ctx *gin.Context) {
	var _ mc.HTTPSuccess
	// Validation
	var req mo.MerchantRequest
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
	// Create
	item, err := q.CreateMerchant(ctx, db.CreateMerchantParams{
		MerchantName: req.MerchantName,
		Commission:   req.Commission,
	})
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error creating merchant",
		})
		return
	}
	// Response
	ctx.JSON(http.StatusCreated, getMerchant(item))
}

// Read godoc
// @Summary Read
// @Description Read
// @Tags Merchants
// @Param   id     path    int32     true        "1"
// @Accept json
// @Produce json
// @Success      200  {object}  mo.Merchant
// @Failure      400  {object}  mc.HTTPError
// @Failure      404  {object}  mc.HTTPError
// @Failure      500  {object}  mc.HTTPError
// @Security Bearer
// @securityDefinitions.apikey Bearer
// @in header
// @name authorization
// @description	Type "Bearer" followed by a space and JWT token.
// @Router /merchant/{id} [get]
func (s *MerchantController) ReadMerchant(ctx *gin.Context) {
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
	// Read
	item, err := q.ReadMerchant(ctx, int32(id))
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error creating merchant",
		})
		return
	}
	// Response
	ctx.JSON(http.StatusCreated, getMerchant(item))
}

// Update godoc
// @Summary Update
// @Description Update
// @Tags Merchants
// @Param   id     path    int32     true        "1"
// @Param   req body mo.MerchantRequest true "Update"
// @Accept json
// @Produce json
// @Success      200  {object}  mo.Merchant
// @Failure      400  {object}  mc.HTTPError
// @Failure      404  {object}  mc.HTTPError
// @Failure      500  {object}  mc.HTTPError
// @Security Bearer
// @securityDefinitions.apikey Bearer
// @in header
// @name authorization
// @description	Type "Bearer" followed by a space and JWT token.
// @Router /merchant/{id} [patch]
func (s *MerchantController) UpdateMerchant(ctx *gin.Context) {
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
	// Validation
	var req mo.MerchantRequest
	err = ctx.ShouldBindJSON(&req)
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
	// Update
	item, err := q.UpdateMerchant(ctx, db.UpdateMerchantParams{
		MerchantID:   int32(id),
		MerchantName: req.MerchantName,
		Commission:   req.Commission,
	})
	if err != nil {
		// Logger
		s.logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Response
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error creating merchant",
		})
		return
	}
	// Response
	ctx.JSON(http.StatusCreated, getMerchant(item))
}
