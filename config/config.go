package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type EnvValues struct {
	UrlPostgresConnection string
	ServerPort            string
}

func LoadEnv() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Could not load .env file X_x.\n")
	}
	fmt.Println(" .env loaded successfully O_o")
}

func GetEnvironment() EnvValues {
	var serverPort = "3000"

	if os.Getenv("SERVER_PORT") != "" {
		serverPort = os.Getenv("SERVER_PORT")
	}

	environmentValues := EnvValues{
		UrlPostgresConnection: os.Getenv("CONNECTION_POSTGRESQL"),
		ServerPort:            serverPort,
	}

	return environmentValues
}
