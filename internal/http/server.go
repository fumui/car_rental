package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HTTPServer struct {
	r  *gin.Engine
	db *pgxpool.Pool
}

func NewHttpServer(db *pgxpool.Pool) *HTTPServer {
	return &HTTPServer{db: db, r: gin.Default()}
}

func (s *HTTPServer) Run() {
	s.r.Run()
}
func (s *HTTPServer) GetRouter() *gin.Engine {
	return s.r
}

func (s *HTTPServer) RegisterRoutes() {
	s.r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s.CreateCar()
	s.GetCars()
	s.GetCarById()
	s.UpdateCar()
	s.DeleteCar()
}
