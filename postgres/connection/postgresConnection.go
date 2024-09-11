package connection

import (
	"context"
	"firstGolang/postgres/config"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type PostgresDatabase struct {
}

func (pgD *PostgresDatabase) Connect() *pgx.Conn {
	ctx := context.Background()

	fmt.Println("Connecting to postgres")

	databaseConfig := config.NewDatabaseConfigPostgres()
	connection, err := pgx.Connect(ctx, databaseConfig.GetUrlPostgres())

	if err != nil {
		fmt.Println("Error postgres ::: ", err)
	} else {
		fmt.Println("Database connected successfully")
	}

	return connection
}

func NewConnectPostgres() *PostgresDatabase {
	fmt.Println("Start postgres connection")
	return &PostgresDatabase{}
}
