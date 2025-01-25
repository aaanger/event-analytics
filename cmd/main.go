package main

import (
	"github.com/aaanger/event-analytics/pkg/db"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %s", err)
	}

	db, err := db.Open(db.ClickHouseConfig{
		Username: os.Getenv("clickhouse_user"),
		Password: os.Getenv("clickhouse_password"),
		Database: os.Getenv("clickhouse_dbname"),
	})

	if err != nil {
		log.Fatalf("error connecting to clickhouse db: %s", err)
	}

	server := NewServer(db, os.Getenv("server_port"))
	err = server.Run()
	if err != nil {
		log.Fatalf("error running grpc server: %s", err)
	}
}
