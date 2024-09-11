package main

import (
	"context"
	"firstGolang/adapters/config"
	"firstGolang/adapters/connection"
	"firstGolang/infrastructure"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func main() {
	err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Println("Error loading .env:", err)
		return
	}

	envConfig := config.NewEnvAdapter()

	// host := envConfig.GetHost()
	serverPort := envConfig.GetServerPort()
	databaseUrl := envConfig.GetPostgrestUrl()

	// databaseUrl := fmt.Sprintf("postgres://%s", host)
	// fmt.Println("serverPort: ", serverPort)
	// fmt.Println("databaseUrl: ", databaseUrl)

	conn, err := connection.ConnectToPostgres(databaseUrl)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			fmt.Println("Error connecting to database:", err)
		}
	}(conn, context.Background())

	infrastructure.StartHTTPServer(serverPort)
}
