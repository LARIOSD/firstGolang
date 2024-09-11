package config

import "os"

type DatabaseConfig struct {
	url string
}

func (config *DatabaseConfig) GetUrlPostgres() string {
	url := os.Getenv("CONNECTION_POSTGRESQL")
	return url
}

func NewDatabaseConfigPostgres() *DatabaseConfig {
	return &DatabaseConfig{}
}
