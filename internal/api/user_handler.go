package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saeede-bellefille/simple-backend/internal/api/dto"
	"github.com/saeede-bellefille/simple-backend/internal/auth"
	"github.com/saeede-bellefille/simple-backend/internal/domain"
	"github.com/saeede-bellefille/simple-backend/internal/middleware"
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

	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
	g.GET("/test", h.Test)

	protected := g.Group("", middleware.JWT())
	protected.GET("/:username", h.Get)
	protected.PUT("/profile", h.UpdateProfile)
	protected.POST("/change-password", h.ChangePassword)

	admin := g.Group("", middleware.JWT(), middleware.RequireRole(domain.RoleAdmin))
	admin.GET("/list", h.List)
	admin.PUT("/:username/role", h.UpdateRole)
}

func (h *userHandler) Get(c echo.Context) error {
	tokenUsername := c.Get("username").(string)
	tokenRole := c.Get("role").(string)
	requestedUsername := c.Param("username")

	if tokenRole != string(domain.RoleAdmin) && tokenUsername != requestedUsername {
		return c.JSON(http.StatusForbidden, &dto.Error{
			Message: "Access denied",
		})
	}

	u, err := h.service.Read(requestedUsername)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, u)
}
func (h *userHandler) List(c echo.Context) error {
	users, err := h.service.List()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *userHandler) Register(c echo.Context) error {
	var data dto.UserRegister
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	user := domain.User{
		Username: data.Username,
		Email:    data.Email,
		Name:     data.Name,
		Age:      data.Age,
		Role:     domain.RoleUser,
	}

	err := h.service.Register(&user, data.Password, data.Repeat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &dto.Error{
		Message: "User registered successfully",
	})
}
func (h *userHandler) Login(c echo.Context) error {
	var data dto.UserLogin
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	u, err := h.service.Login(data.Username, data.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &dto.Error{
			Message: "Invalid credentials",
		})
	}

	token, err := auth.GenerateToken(u.Username, string(u.Role))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &dto.Error{
			Message: "Could not generate token",
		})
	}

	response := dto.LoginResponse{
		Token: token,
		User: dto.User{
			Username: u.Username,
			Email:    u.Email,
			Name:     u.Name,
			Age:      u.Age,
			Role:     string(u.Role),
		},
	}

	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateProfile(c echo.Context) error {
	username := c.Get("username").(string)

	var data dto.User
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	user := domain.User{
		Username: username,
		Email:    data.Email,
		Name:     data.Name,
		Age:      data.Age,
	}

	err := h.service.UpdateProfile(username, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &dto.Error{
		Message: "Profile updated successfully",
	})
}

func (h *userHandler) UpdateRole(c echo.Context) error {
	username := c.Param("username")

	var data dto.UserRole
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	role := domain.Role(data.Role)
	err := h.service.UpdateRole(username, role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &dto.Error{
		Message: "Role updated successfully",
	})
}

func (h *userHandler) ChangePassword(c echo.Context) error {
	username := c.Get("username").(string)

	var data struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
		RepeatPassword  string `json:"repeat_password"`
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	err := h.service.ChangePassword(username, data.CurrentPassword, data.NewPassword, data.RepeatPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &dto.Error{
		Message: "Password changed successfully",
	})
}

func (h *userHandler) Test(c echo.Context) error {
	return c.String(http.StatusOK, h.service.Test())
}
