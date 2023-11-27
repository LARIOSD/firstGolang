package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"os"
)

func PostgresConnection() {
	urlConnection := "postgresql://cerbero:JkrqUP4N6Tn3cB2qXKNinLl4zo61gUZTJu5qCsN2J94wNa6QC6@144.91.124.50:12015/hades"

	conn, err := pgx.Connect(context.Background(), urlConnection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	print("Database connection successfully !\n")
	defer conn.Close(context.Background())

	var date pgtype.Date
	conn.QueryRow(context.Background(), "select now() as date").Scan(&date)
	fmt.Println(date)
}
