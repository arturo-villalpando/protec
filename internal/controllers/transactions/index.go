package transactions

import (
	"protec/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type TransactionController struct {
	config config.ServerConfig
	dbConn *pgxpool.Pool
	logger *zerolog.Logger
}

func NewTransactionController(conf config.ServerConfig, dbConn *pgxpool.Pool, logger *zerolog.Logger) *TransactionController {
	cLogger := logger.With().Str("controller", "Transactions controller logger").Logger()
	server := &TransactionController{
		config: conf,
		dbConn: dbConn,
		logger: &cLogger,
	}

	return server
}
