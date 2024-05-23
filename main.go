package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"protec/config"
	"protec/docs"

	cm "protec/internal/controllers/merchants"
	rm "protec/internal/routers/merchants"

	ct "protec/internal/controllers/transactions"
	rt "protec/internal/routers/transactions"

	cr "protec/internal/controllers/reports"
	rr "protec/internal/routers/reports"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	ctx context.Context

	server *gin.Engine

	c       config.Config
	srvConf config.ServerConfig

	logger *zerolog.Logger

	MerchantController cm.MerchantController
	MerchantRouter     rm.MerchantRouter

	TransactionController ct.TransactionController
	TransactionRouter     rt.TransactionRouter

	ReportController cr.ReportController
	ReportRouter     rr.ReportRouter
)

func init() {
	ctx = context.Background()
	// Conf vars
	c = config.ViperConfigInit()
	srvConf = *c.GetServerConfig()
	// Server
	if srvConf.ENV != "local" {
		gin.SetMode(gin.ReleaseMode)
	}
	server = gin.Default()
	// Logger
	logger = config.Logger()
}

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Programmatically set swagger info
	docs.SwaggerInfo.Title = "Protec API"
	docs.SwaggerInfo.Description = "This is the protec documentation"
	docs.SwaggerInfo.Version = srvConf.VERSION
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// Database connection
	dbConn, err := config.DBConnect(ctx, srvConf.DB_CONNECTION)
	if err != nil {
		// Logger
		logger.Error().
			Str("Error", err.Error()).
			Str("Code", fmt.Sprint(http.StatusInternalServerError))
		// Fatal
		log.Fatalf("db connection error: %v", err)
	}
	defer dbConn.Close()
	// CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-Requested-With"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	// Fix error with cors and group routes more info:
	// https://github.com/gin-gonic/gin/issues/3546
	// server.OPTIONS("/*api")
	// Create a healthcheck endpoint and keep hide
	server.GET("/healthcheck/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "pong",
		})
	})
	// Routes
	MerchantController = *cm.NewMerchantController(srvConf, dbConn, logger)
	MerchantRouter = *rm.NewMerchantRouter(MerchantController, srvConf, logger)

	MerchantRouter.MerchantRouter(server)

	TransactionController = *ct.NewTransactionController(srvConf, dbConn, logger)
	TransactionRouter = *rt.NewTransactionRouter(TransactionController, srvConf, logger)

	TransactionRouter.TransactionRouter(server)

	ReportController = *cr.NewReportController(srvConf, dbConn, logger)
	ReportRouter = *rr.NewReportRouter(ReportController, srvConf, logger)

	ReportRouter.ReportTransactionRouter(server)
	// Add swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Lounch error if the routs doesnt exists
	server.NoRoute(func(ctx *gin.Context) {
		// Logger
		logger.Error().
			Str("Error", fmt.Sprintf("Route %s not found", ctx.Request.URL)).
			Str("Code", fmt.Sprint(http.StatusNotFound))
		// Fatal
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	// Run gin
	log.Fatal(server.Run(":3000"))
}
