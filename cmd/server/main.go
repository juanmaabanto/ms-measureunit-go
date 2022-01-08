package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/ports"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/service"
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

	router.Use(
	// middleware.Recover(), // Recover from all panics to always have your server up
	// middleware.Logger(),    // Log everything to stdout
	// middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)

	// router.HTTPErrorHandler = func(err error, c echo.Context) {
	// 	// Take required information from error and context and send it to a service like New Relic
	// 	fmt.Println(c.Path(), c.QueryParams(), err.Error())

	// 	// Call the default handler to return the HTTP response
	// 	router.DefaultHTTPErrorHandler(err, c)
	// }

	router.POST("/api/v1/measuretypes", si.CreateMeasureType)
}
