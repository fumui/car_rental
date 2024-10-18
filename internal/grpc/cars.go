package grpc

import (
	pb "carrental/api"
	"carrental/pkg/cars"
	"context"
)

func (s *GRPCServer) CreateCar(ctx context.Context, req *pb.Car) (*pb.Car, error) {
	return cars.CreateCar(ctx, s.db, req)
}

func (s *GRPCServer) GetCars(ctx context.Context, req *pb.GetCarsRequest) (*pb.GetCarsResponse, error) {
	return cars.GetCars(ctx, s.db, req)
}

func (s *GRPCServer) GetCarById(ctx context.Context, req *pb.Car) (*pb.Car, error) {
	return cars.GetCarById(ctx, s.db, req)
}

func (s *GRPCServer) UpdateCar(ctx context.Context, req *pb.Car) (*pb.Car, error) {
	return cars.UpdateCar(ctx, s.db, req)
}

func (s *GRPCServer) DeleteCar(ctx context.Context, req *pb.Car) (*pb.Car, error) {
	return cars.DeleteCar(ctx, s.db, req)
}
