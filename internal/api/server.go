package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/saeede-bellefille/simple-backend/internal/repository/models"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &Server{echo: e}
}

func (s *Server) Setup(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Product{})

	setupUser(s.echo.Group("/user"), db)
	setupProduct(s.echo.Group("/product"), db)
}

func (s *Server) Run(address string) error {
	return s.echo.Start(address)
}
