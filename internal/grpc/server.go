package grpc

import (
	pb "carrental/api"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GRPCServer struct {
	pb.UnimplementedCarRentalServicesServer
	db *pgxpool.Pool
}

func NewGRPCServer(db *pgxpool.Pool) *GRPCServer {
	return &GRPCServer{db: db}
}
