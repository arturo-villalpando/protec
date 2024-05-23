package config

type ServerConfig struct {
	// Enviroment
	ENV     string
	VERSION string
	// Host
	HOST    string
	ORIGINS string
	// Database
	DB_CONNECTION string
}

type config struct {
	serverConfig *ServerConfig
}

type Config interface {
	GetConfig() *config
	GetServerConfig() *ServerConfig
}
