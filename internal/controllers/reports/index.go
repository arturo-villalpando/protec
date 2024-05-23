package reports

import (
	"protec/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type ReportController struct {
	config config.ServerConfig
	dbConn *pgxpool.Pool
	logger *zerolog.Logger
}

func NewReportController(conf config.ServerConfig, dbConn *pgxpool.Pool, logger *zerolog.Logger) *ReportController {
	cLogger := logger.With().Str("controller", "Reports controller logger").Logger()
	server := &ReportController{
		config: conf,
		dbConn: dbConn,
		logger: &cLogger,
	}

	return server
}
