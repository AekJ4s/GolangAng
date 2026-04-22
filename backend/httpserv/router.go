package httpserv

import (
	"strings"

	"golangbackend/infrastructure"
	"golangbackend/internal/adaptor/handler"
	"golangbackend/internal/adaptor/repo"
	"golangbackend/internal/core/service"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	userRepo := repo.NewUserRepo(infrastructure.DB)
	userSvc := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	api := e.Group("/api")
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)

	protected := api.Group("", jwtMiddleware)
	protected.GET("/profile", userHandler.Profile)
}

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(401, map[string]string{"message": "unauthorized"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if _, err := service.ValidateToken(tokenString); err != nil {
			return c.JSON(401, map[string]string{"message": "unauthorized"})
		}

		return next(c)
	}
}
