package cars

import (
	"context"
	"fmt"
	"log"

	pb "carrental/api"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateCar(ctx context.Context, db *pgxpool.Pool, req *pb.Car) (*pb.Car, error) {
	if req.Name == "" {
		return nil, ErrCarNameIsRequired
	}
	err := db.QueryRow(ctx, `
		INSERT INTO cars (name,day_rate,month_rate,image) 
		VALUES ($1, $2, $3, $4)
		RETURNING id`,
		req.Name,
		req.DayRate,
		req.MonthRate,
		req.Image,
	).Scan(&req.Id)
	if err != nil {
		log.Println("error on creating car: ", err)
		return nil, err
	}
	return req, nil
}

func GetCars(ctx context.Context, db *pgxpool.Pool, req *pb.GetCarsRequest) (*pb.GetCarsResponse, error) {
	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	query := `
	SELECT
		cars.id,
		cars.name,
		cars.day_rate,
		cars.month_rate,
		cars.image
	FROM cars `
	params := []interface{}{req.Limit, (req.Page * req.Limit) - req.Limit}
	if req.Search != "" {
		params = append(params, "%"+req.Search+"%")
		query += " WHERE cars.name ILIKE $" + fmt.Sprint(len(params))
	}
	query += " LIMIT $1 OFFSET $2"

	rows, err := db.Query(ctx, query, params...)
	if err != nil {
		log.Println("error on getting list of cars: ", err)
		return nil, err
	}
	defer rows.Close()
	res := &pb.GetCarsResponse{
		Limit: req.Limit,
		Page:  req.Page,
	}
	for rows.Next() {
		car := &pb.Car{}
		err := rows.Scan(
			&car.Id,
			&car.Name,
			&car.DayRate,
			&car.MonthRate,
			&car.Image,
		)
		if err != nil {
			log.Println("error on getting list of cars: ", err)
			return nil, err
		}
		res.Data = append(res.Data, car)
	}
	return res, nil
}

func GetCarById(ctx context.Context, db *pgxpool.Pool, req *pb.Car) (*pb.Car, error) {
	if req.Id == 0 {
		return nil, ErrCarIdIsRequired
	}
	query := `
	SELECT
		cars.id,
		cars.name,
		cars.day_rate,
		cars.month_rate,
		cars.image
	FROM cars
	WHERE cars.id = $1`
	err := db.QueryRow(ctx, query, req.Id).Scan(
		&req.Id,
		&req.Name,
		&req.DayRate,
		&req.MonthRate,
		&req.Image,
	)
	if err != nil {
		log.Printf("error on getting car id (%d): %v", req.Id, err)
		return nil, err
	}
	return req, nil
}
func UpdateCar(ctx context.Context, db *pgxpool.Pool, req *pb.Car) (*pb.Car, error) {
	if req.Id == 0 {
		return nil, ErrCarIdIsRequired
	}
	if req.Name == "" {
		return nil, ErrCarNameIsRequired
	}
	query := `UPDATE cars
	SET name = $1, day_rate = $2, month_rate = $3, image = $4
	WHERE id = $5`
	tag, err := db.Exec(ctx, query, req.Name, req.DayRate, req.MonthRate, req.Image, req.Id)
	if err != nil {
		log.Printf("error on updating car id (%d): %v", req.Id, err)
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		return nil, ErrCarNotFound
	}
	return req, nil
}
func DeleteCar(ctx context.Context, db *pgxpool.Pool, req *pb.Car) (*pb.Car, error) {
	if req.Id == 0 {
		return nil, ErrCarIdIsRequired
	}
	query := `DELETE FROM cars WHERE id = $1`
	tag, err := db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Printf("error on updating car id (%d): %v", req.Id, err)
		return nil, err
	}
	if tag.RowsAffected() == 0 {
		return nil, ErrCarNotFound
	}
	return req, nil
}
