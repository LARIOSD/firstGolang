package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type PostgresTimeAdapter struct {
	conn *pgx.Conn
}

func NewPostgresTimeAdapter(conn *pgx.Conn) *PostgresTimeAdapter {
	return &PostgresTimeAdapter{
		conn: conn,
	}
}

func (adapter *PostgresTimeAdapter) GetCurrentTime() (string, error) {
	ctx := context.Background()
	var currentTime string

	err := adapter.conn.QueryRow(ctx, "SELECT NOW()").Scan(&currentTime)
	if err != nil {
		return "", fmt.Errorf("error executing query: %w", err)
	}

	return currentTime, nil
}
