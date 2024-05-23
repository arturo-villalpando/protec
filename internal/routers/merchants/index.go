package merchants

import (
	"protec/config"
	co "protec/internal/controllers/merchants"

	"github.com/rs/zerolog"
)

type MerchantRouter struct {
	controller co.MerchantController
	config     config.ServerConfig
	logger     *zerolog.Logger
}

func NewMerchantRouter(controller co.MerchantController, conf config.ServerConfig, logger *zerolog.Logger) *MerchantRouter {
	rLogger := logger.With().Str("router", "Merchant router logger").Logger()
	server := &MerchantRouter{
		controller: controller,
		config:     conf,
		logger:     &rLogger,
	}

	return server

}
