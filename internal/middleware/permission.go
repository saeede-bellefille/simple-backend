package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saeede-bellefille/simple-backend/internal/api/dto"
	"github.com/saeede-bellefille/simple-backend/internal/domain"
)

func RequireRole(requiredRoles ...domain.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Get("role")
			if userRole == nil {
				return c.JSON(http.StatusUnauthorized, &dto.Error{
					Message: "Authentication required",
				})
			}

			role := domain.Role(userRole.(string))
			for _, requiredRole := range requiredRoles {
				if hasPermission(role, requiredRole) {
					return next(c)
				}
			}

			return c.JSON(http.StatusForbidden, &dto.Error{
				Message: "Insufficient permissions",
			})
		}
	}
}

func hasPermission(userRole, requiredRole domain.Role) bool {
	if userRole == domain.RoleAdmin {
		return true
	}

	if userRole == domain.RoleModerator && requiredRole == domain.RoleUser {
		return true
	}

	return userRole == requiredRole
}
