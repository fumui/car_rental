package main

import (
	http "carrental/internal/http"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("C:\\Users\\fuadm\\Work\\personal\\template\\.env")
	if err != nil {
		log.Fatal(err)
	}
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	httpServer := http.NewHttpServer(dbpool)
	httpServer.RegisterRoutes()
	httpServer.Run()
	log.Println("Hello, World!")
}
