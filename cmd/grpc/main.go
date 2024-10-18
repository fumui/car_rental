package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "carrental/api"

	impl "carrental/internal/grpc"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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
	s := grpc.NewServer()
	implementor := impl.NewGRPCServer(dbpool)
	pb.RegisterCarRentalServicesServer(s, implementor)
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
