package main

import (
	"context"
	"log"

	_ "github.com/sofisoft-tech/ms-measureunit/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/ports"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/service"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/validations"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/managers"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Measure Unit API
// @version v1
// @description Specifying services for measure units.

// @contact.name Sofisoft Technologies SAC
// @contact.url https://sofisoft.pe
// @contact.email juan.abanto@sofisoft.pe

// @license.name MIT License
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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

	//Swagger
	router.GET("/*", echoSwagger.WrapHandler)

	//measureType
	api.POST("/measuretypes", si.CreateMeasureType)
}
