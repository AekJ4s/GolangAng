package infrastructure

import (
	"strings"

	"github.com/labstack/echo/v4"
)

var allowedOrigins = []string{
	"http://localhost:4200",
	"http://localhost:80",
	"http://localhost",
}

func InitMiddleware(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get("Origin")
			for _, allowed := range allowedOrigins {
				if strings.EqualFold(origin, allowed) {
					c.Response().Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}
			return next(c)
		}
	})
}
