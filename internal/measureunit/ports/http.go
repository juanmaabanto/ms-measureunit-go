package ports

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app/command"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

func (h HttpServer) CreateMeasureType(c echo.Context) error {
	item := command.CreateMeasureType{}

	if err := c.Bind(&item); err != nil {
		panic(err)
	}

	id, err := h.app.Commands.CreateMeasureType.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	// create a response
	response := map[string]interface{}{"id": id}
	//return success response
	return c.JSON(http.StatusCreated, response)
}
