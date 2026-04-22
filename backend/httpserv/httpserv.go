package httpserv

import (
	"golangbackend/infrastructure"

	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	e.HideBanner = true

	infrastructure.InitMiddleware(e)

	RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":" + infrastructure.AppCfg.App.Port))
}
