package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/ports"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/service"
)

var (
	router = echo.New()
)

func main() {
	ctx := context.Background()

	application := service.NewApplication(ctx)

	Handler(ports.NewHttpServer(application), router)
	router.Logger.Fatal(router.Start(":3000"))
}

type ServerInterface interface {
	CreateMeasureType(c echo.Context) error
}

func Handler(si ServerInterface, router *echo.Echo) {
	if router == nil {
		router = echo.New()
	}

	router.POST("/api/v1/measuretypes", si.CreateMeasureType)
}
