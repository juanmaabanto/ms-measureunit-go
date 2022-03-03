package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

func main() {
	// create a Dapr service (e.g. ":8080", "0.0.0.0:8080", "10.1.1.1:8080" )
	s := daprd.NewService(":8080")

	// add some topic subscriptions
	sub := &common.Subscription{
		PubsubName: "messages",
		Topic:      "topic1",
		Route:      "/events",
	}
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	// add a service to service invocation handler
	if err := s.AddServiceInvocationHandler("/echo", echoHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	// add an input binding invocation handler
	if err := s.AddBindingInvocationHandler("/run", runHandler); err != nil {
		log.Fatalf("error adding binding handler: %v", err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName:%s, Topic:%s, ID:%s, Data: %s", e.PubsubName, e.Topic, e.ID, e.Data)
	return false, nil
}

func echoHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}

	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}

func runHandler(ctx context.Context, in *common.BindingEvent) (out []byte, err error) {
	log.Printf("binding - Data:%s, Meta:%v", in.Data, in.Metadata)
	return nil, nil
}

// package main

// import (
// 	"context"
// 	"log"

// 	_ "github.com/sofisoft-tech/ms-measureunit/docs"

// 	"github.com/joho/godotenv"
// 	"github.com/labstack/echo/v4"
// 	"github.com/sofisoft-tech/ms-measureunit/internal/ports"
// 	"github.com/sofisoft-tech/ms-measureunit/internal/service"
// 	"github.com/sofisoft-tech/ms-measureunit/internal/validations"
// 	"github.com/sofisoft-tech/ms-measureunit/seedwork/managers"
// 	"github.com/sofisoft-tech/ms-measureunit/seedwork/middleware"
// 	echoSwagger "github.com/swaggo/echo-swagger"
// )

// // @title MeasureUnit API
// // @version v1
// // @description Specifying services for measure units.

// // @contact.name Sofisoft Technologies SAC
// // @contact.url https://sofisoft.pe
// // @contact.email juan.abanto@sofisoft.pe

// // @license.name MIT License
// // @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	router := echo.New()
// 	ctx := context.Background()

// 	application := service.NewApplication(ctx)

// 	Handler(ports.NewHttpServer(application), router)
// 	router.Logger.Fatal(router.Start(":3000"))
// }

// type ServerInterface interface {
// 	AddMeasureType(c echo.Context) error
// 	DeleteMeasureType(c echo.Context) error
// 	GetMeasureType(c echo.Context) error
// 	ListMeasureType(c echo.Context) error
// 	UpdateMeasureType(c echo.Context) error
// }

// func Handler(si ServerInterface, router *echo.Echo) {
// 	if router == nil {
// 		router = echo.New()
// 	}

// 	router.Validator = validations.NewValidationUtil()

// 	loggerManager := managers.NewLoggerManager("https://services.sofisoft.pe/logging/", "ms-measureunit")

// 	api := router.Group("/api/v1")

// 	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 		LoggerErrorFunc: loggerManager.Error,
// 	}))

// 	//Swagger
// 	router.GET("/*", echoSwagger.WrapHandler)

// 	//measureType
// 	api.DELETE("/measuretypes/:id", si.DeleteMeasureType)
// 	api.GET("/measuretypes", si.ListMeasureType)
// 	api.GET("/measuretypes/:id", si.GetMeasureType)
// 	api.PATCH("/measuretypes", si.UpdateMeasureType)
// 	api.POST("/measuretypes", si.AddMeasureType)
// }
