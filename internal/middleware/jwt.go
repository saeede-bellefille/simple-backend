package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/saeede-bellefille/simple-backend/internal/api/dto"
	"github.com/saeede-bellefille/simple-backend/internal/auth"
)

func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHandler := c.Request().Header.Get("Authorization")
			if authHandler == "" {
				return c.JSON(http.StatusUnauthorized, &dto.Error{
					Message: "Authorization header required",
				})
			}
			if !strings.HasPrefix(authHandler, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, &dto.Error{
					Message: "Bearer token required",
				})
			}
			tokenString := strings.TrimPrefix(authHandler, "Bearer ")
			claims, err := auth.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, &dto.Error{
					Message: "Invalid token",
				})
			}
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
			return next(c)
		}
	}

}
