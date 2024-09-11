package config

import (
	"os"
)

type EnvAdapter struct{}

func NewEnvAdapter() *EnvAdapter {
	return &EnvAdapter{}
}

func (e *EnvAdapter) GetHost() string {
	return os.Getenv("HOST")
}

func (e *EnvAdapter) GetPostgrestUrl() string {
	return os.Getenv("CONNECTION_POSTGRESQL")
}

func (e *EnvAdapter) GetServerPort() string {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "3001"
	}
	return serverPort
}
