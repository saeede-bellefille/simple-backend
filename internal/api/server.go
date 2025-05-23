package api

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	return &Server{echo: echo.New()}
}

func (s *Server) Setup(db *gorm.DB) {
	setupUser(s.echo.Group("/user"), db)
}

func (s *Server) Run(address string) error {
	return s.echo.Start(address)
}
