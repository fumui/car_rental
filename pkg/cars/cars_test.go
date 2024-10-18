package cars

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	pb "carrental/api"
)

var dbConnPool *pgxpool.Pool

func PrepareDB() {
	if dbConnPool != nil {
		return
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	migrationPath := os.Getenv("MIGRATION_PATH")
	if migrationPath == "" {
		log.Println("MIGRATION_PATH environment is not set, using default value")
		cwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		migrationPath = filepath.Join(cwd, "migrations", "test")
	}

	// Convert backslashes to forward slashes for Windows compatibility
	migrationPath = strings.ReplaceAll(migrationPath, "\\", "/")
	migrationPath = "file://" + migrationPath

	url := os.Getenv("DATABASE_URL")
	log.Println(fmt.Sprintf("Migrating %s from %s", url, migrationPath))
	m, err := migrate.New(migrationPath, url)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Drop()
	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	m, err = migrate.New(migrationPath, url)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	dbConnPool = dbpool
}

func TestCreateCar(t *testing.T) {
	PrepareDB()
	car := &pb.Car{
		Name:      "BMW",
		DayRate:   100000,
		MonthRate: 2000000,
		Image:     "bmw.png",
	}

	car, err := CreateCar(context.Background(), dbConnPool, car)
	if err != nil {
		t.Error("expecting success creating car")
	}
	if car == nil {
		t.Error("expecting car created")
		return
	}
	if car.Id == 0 {
		t.Error("expecting id autofilled")
	}
}

func TestGetCars(t *testing.T) {
	PrepareDB()
	req := &pb.GetCarsRequest{}

	cars, err := GetCars(context.Background(), dbConnPool, req)
	if err != nil || cars == nil {
		t.Error("expecting success get car list")
		return
	}
	if cars.Limit != 10 {
		t.Error("expecting default limit 10")
	}
	if cars.Page != 1 {
		t.Error("expecting default page 1")
	}
	println(cars)
	if len(cars.Data) != 1 {
		t.Errorf("expecting 1 car data but got %d", len(cars.Data))
	}
}

func TestGetCarById(t *testing.T) {
	PrepareDB()
	req := &pb.Car{Id: 1}

	car, err := GetCarById(context.Background(), dbConnPool, req)
	if err != nil || car == nil {
		t.Error("expecting success creating car")
		return
	}
	if car.Id != 1 {
		t.Error("expecting car id 1")
	}
	if !strings.EqualFold(car.Name, req.Name) {
		t.Errorf("expecting first car to be BMW, but got %s", car.Name)
	}
}

func TestUpdateCar(t *testing.T) {
	PrepareDB()
	req := &pb.Car{
		Id:        1,
		Name:      "Lamborghini",
		DayRate:   200000,
		MonthRate: 4000000,
		Image:     "Lamborghini.png",
	}

	car, err := UpdateCar(context.Background(), dbConnPool, req)
	if err != nil || car == nil {
		t.Error("expecting success updating car")
		return
	}
	if car.Id != 1 {
		t.Error("expecting car id 1")
	}
	if !strings.EqualFold(car.Name, req.Name) {
		t.Errorf("expecting first car to be %s", req.Name)
	}
}

func TestDeleteCar(t *testing.T) {
	PrepareDB()
	req := &pb.Car{
		Id: 1,
	}
	car, err := DeleteCar(context.Background(), dbConnPool, req)
	if err != nil || car == nil {
		t.Error("expecting success deleting car")
		return
	}
	if car.Id != 1 {
		t.Error("expecting car id 1")
	}
}
