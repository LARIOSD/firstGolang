package ports

type ConfigPort interface {
	GetHost() string
	GetServerPort() string
	GetPostgrestUrl() string
}
