package main

import (
	"context"
	"log"

	_ "github.com/sofisoft-tech/ms-measureunit/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/ports"
	"github.com/sofisoft-tech/ms-measureunit/internal/service"
	"github.com/sofisoft-tech/ms-measureunit/internal/validations"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/managers"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title MeasureUnit API
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
	AddMeasureType(c echo.Context) error
	GetMeasureType(c echo.Context) error
	ListMeasureType(c echo.Context) error
	UpdateMeasureType(c echo.Context) error
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
	api.GET("/measuretypes", si.ListMeasureType)
	api.GET("/measuretypes/:id", si.GetMeasureType)
	api.PATCH("/measuretypes", si.UpdateMeasureType)
	api.POST("/measuretypes", si.AddMeasureType)
}
