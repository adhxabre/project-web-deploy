package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	var err error

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Success connect to database")

}
