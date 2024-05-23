package merchants

import (
	"protec/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type MerchantController struct {
	config config.ServerConfig
	dbConn *pgxpool.Pool
	logger *zerolog.Logger
}

func NewMerchantController(conf config.ServerConfig, dbConn *pgxpool.Pool, logger *zerolog.Logger) *MerchantController {
	cLogger := logger.With().Str("controller", "Merchants controller logger").Logger()
	server := &MerchantController{
		config: conf,
		dbConn: dbConn,
		logger: &cLogger,
	}

	return server
}
