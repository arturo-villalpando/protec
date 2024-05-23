package reports

import (
	"protec/config"
	co "protec/internal/controllers/reports"

	"github.com/rs/zerolog"
)

type ReportRouter struct {
	controller co.ReportController
	config     config.ServerConfig
	logger     *zerolog.Logger
}

func NewReportRouter(controller co.ReportController, conf config.ServerConfig, logger *zerolog.Logger) *ReportRouter {
	rLogger := logger.With().Str("router", "Report router logger").Logger()
	server := &ReportRouter{
		controller: controller,
		config:     conf,
		logger:     &rLogger,
	}

	return server

}
