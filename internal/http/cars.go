package http

import (
	pb "carrental/api"
	"carrental/pkg/cars"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s HTTPServer) CreateCar() {
	s.r.POST("/cars", func(ctx *gin.Context) {
		req := &pb.Car{}
		err := ctx.ShouldBindBodyWithJSON(req)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "invalid request",
			})
			return
		}

		req, err = cars.CreateCar(ctx.Copy(), s.db, req)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, req)
	})
}

func (s HTTPServer) GetCars() {
	s.r.GET("/cars", func(ctx *gin.Context) {
		req := &pb.GetCarsRequest{}
		err := ctx.ShouldBindBodyWithJSON(req)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "invalid request",
			})
			return
		}

		res, err := cars.GetCars(ctx.Copy(), s.db, req)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, res)
	})
}

func (s HTTPServer) GetCarById() {
	s.r.GET("/cars/:id", func(ctx *gin.Context) {
		param, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(400, map[string]interface{}{
				"message": "invalid request",
			})
			return
		}
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "invalid parameter id",
			})
			return
		}
		car := &pb.Car{
			Id: id,
		}

		res, err := cars.GetCarById(ctx.Copy(), s.db, car)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, res)
	})
}

func (s HTTPServer) UpdateCar() {
	s.r.PUT("/cars", func(ctx *gin.Context) {
		req := &pb.Car{}
		err := ctx.ShouldBindBodyWithJSON(req)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "invalid request",
			})
			return
		}

		res, err := cars.UpdateCar(ctx.Copy(), s.db, req)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, res)
	})
}

func (s HTTPServer) DeleteCar() {
	s.r.DELETE("/cars/:id", func(ctx *gin.Context) {
		param, ok := ctx.Params.Get("id")
		if !ok {
			ctx.JSON(400, map[string]interface{}{
				"message": "invalid request",
			})
			return
		}
		id, err := strconv.ParseInt(param, 10, 64)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "invalid parameter id",
			})
			return
		}
		car := &pb.Car{
			Id: id,
		}

		res, err := cars.DeleteCar(ctx.Copy(), s.db, car)
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, res)
	})
}
