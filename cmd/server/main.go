package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/ports"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/service"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/validations"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/managers"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := echo.New()
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

	router.Validator = validations.NewValidationUtil()

	loggerManager := managers.NewLoggerManager("https://services.sofisoft.pe/logging/", "ms-measureunit")

	api := router.Group("/api/v1")

	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		LoggerErrorFunc: loggerManager.Error,
	}))

	api.POST("/measuretypes", si.CreateMeasureType)
}
