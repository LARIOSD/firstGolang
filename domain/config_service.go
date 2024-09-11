package domain

type ConfigService interface {
	GetHost() string
	GetServerPort() string
	GetPostgrestUrl() string
}
