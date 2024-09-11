package connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type PostgresAdapter struct {
	conn *pgx.Conn
}

func NewPostgresAdapter(conn *pgx.Conn) *PostgresAdapter {
	return &PostgresAdapter{
		conn: conn,
	}
}

func ConnectToPostgres(dbUrl string) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	fmt.Println("Database connected successfully")
	return conn, nil
}
