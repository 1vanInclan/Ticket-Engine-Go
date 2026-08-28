package middleware

import (
	"net/http"
	"strings"
	"ticket-engine/infrastructure/security"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authorization header required"})
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired token"})
			}

			claims, err := security.ValidateToken(parts[1])
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired token"})
			}

			// Guardar claims en el contexto de Echo para los handlers
			c.Set("userID", claims.UserID)
			c.Set("userRole", claims.Role)

			return next(c)
		}
	}
}
