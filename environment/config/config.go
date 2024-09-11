package config

import (
	"os"
)

type EnvValues struct {
	Host       string
	ServerPort string
}

func GetPort() string {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "3001"
	}
	return serverPort
}

func GetEnvironment() EnvValues {
	environmentValues := EnvValues{
		Host:       os.Getenv("HOST"),
		ServerPort: GetPort(),
	}

	return environmentValues
}
