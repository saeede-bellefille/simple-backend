package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saeede-bellefille/simple-backend/internal/api/dto"
	"github.com/saeede-bellefille/simple-backend/internal/domain"
	"github.com/saeede-bellefille/simple-backend/internal/repository"
	"github.com/saeede-bellefille/simple-backend/internal/service/user"
	"gorm.io/gorm"
)

type userHandler struct {
	service *user.Service
}

func setupUser(g *echo.Group, db *gorm.DB) {
	repo := repository.NewUserRepo(db)
	service := user.New(repo)
	h := &userHandler{service: service}
	g.GET("/:username", h.Get)
	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
	g.GET("/test", h.Test)
}

func (h *userHandler) Get(c echo.Context) error {
	u, err := h.service.Read(c.Param("username"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func (h *userHandler) Register(c echo.Context) error {
	var data dto.RegisdterUser
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	user := domain.User{
		Username: data.Username,
		Email:    data.Email,
		Name:     data.Name,
		Age:      data.Age,
	}
	err := h.service.Register(&user, data.Password, data.Repeat)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *userHandler) Login(c echo.Context) error {
	var data dto.LoginUser
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	u, err := h.service.Login(data.Username, data.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func (h *userHandler) Test(c echo.Context) error {
	return c.String(http.StatusOK, h.service.Test())
}
