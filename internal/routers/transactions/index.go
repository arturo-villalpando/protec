package transactions

import (
	"protec/config"
	co "protec/internal/controllers/transactions"

	"github.com/rs/zerolog"
)

type TransactionRouter struct {
	controller co.TransactionController
	config     config.ServerConfig
	logger     *zerolog.Logger
}

func NewTransactionRouter(controller co.TransactionController, conf config.ServerConfig, logger *zerolog.Logger) *TransactionRouter {
	rLogger := logger.With().Str("router", "Transaction router logger").Logger()
	server := &TransactionRouter{
		controller: controller,
		config:     conf,
		logger:     &rLogger,
	}

	return server

}
