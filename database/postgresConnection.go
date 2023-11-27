package database

import (
	"context"
	"firstGolang/config"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"os"
)

func PostgresConnection() {
	environment := config.GetEnvironment()

	conn, err := pgx.Connect(context.Background(), environment.UrlPostgresConnection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	println("Database connection successfully !")

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			fmt.Println("Ups Connection failed X_x")
		}
	}(conn, context.Background())

	var date pgtype.Date
	conn.QueryRow(context.Background(), "select now() as date").Scan(&date)
	fmt.Println(date)
}
