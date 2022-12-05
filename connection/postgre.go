package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	var err error

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	var DB_HOST = os.Getenv("DB_HOST")
	var DB_USER = os.Getenv("DB_USER")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_PORT = os.Getenv("DB_PORT")
	var DB_NAME = os.Getenv("DB_NAME")

	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	check := fmt.Sprintf("%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	fmt.Println(check)
	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Success connect to database")

}
