package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/saeede-bellefille/simple-backend/internal/api/dto"
	"github.com/saeede-bellefille/simple-backend/internal/domain"
	"github.com/saeede-bellefille/simple-backend/internal/repository"
	"github.com/saeede-bellefille/simple-backend/internal/service/product"
	"gorm.io/gorm"
)

type productHandler struct {
	service *product.Service
}

func setupProduct(g *echo.Group, db *gorm.DB) {
	repo := repository.NewProductRepo(db)
	service := product.New(repo)
	h := &productHandler{service: service}
	g.GET("/:id", h.Get)
	g.POST("/create", h.Create)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func (h *productHandler) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64) // Parse as uint64
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: "Invalid ID format (must be a positive number)",
		})
	}
	uintID := uint(id)
	p, err := h.service.Get(uintID)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (h *productHandler) List(c echo.Context) error {
	p, err := h.service.List()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, p)
}

func (h *productHandler) Create(c echo.Context) error {
	var data dto.Product
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	product := domain.Product{
		Name:  data.Name,
		Price: data.Price,
		Group: data.Group,
	}
	err := h.service.Create(&product)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *productHandler) Update(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var data dto.Product
	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	product := domain.Product{
		Name:  data.Name,
		Price: data.Price,
		Group: data.Group,
	}
	err = h.service.Update(uint(id), &product)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *productHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err = h.service.Delete(uint(id)); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
