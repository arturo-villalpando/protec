package config

import (
	"github.com/spf13/viper"
)

type ViperConfig struct {
	config *config
}

func ViperConfigInit() Config {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	// Ignore .env file vars over system env vars
	viper.AutomaticEnv()
	// Read config file
	_ = viper.ReadInConfig()
	return ViperConfig{config: &config{
		serverConfig: &ServerConfig{
			// Enviroment
			ENV:     viper.GetString("ENV"),
			VERSION: viper.GetString("VERSION"),
			// Client Origins
			HOST:    viper.GetString("SERVER_HOST"),
			ORIGINS: viper.GetString("CLIENT_ORIGINS"),
			// Database
			DB_CONNECTION: viper.GetString("DB_CONNECTION"),
		},
	}}
}

func (v ViperConfig) GetServerConfig() *ServerConfig {
	return v.config.serverConfig
}

func (v ViperConfig) GetConfig() *config {
	return v.config
}
