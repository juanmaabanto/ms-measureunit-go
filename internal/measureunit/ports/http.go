package ports

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app/command"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}

// CreateTodo godoc
// @Summary Create a new type of measure.
// @Tags MeasureTypes
// @Accept json
// @Produce json
// @Param command body measuretype.MeasureType true "Object to be created."
// @Success 201 {string} string "Id of the created object"
// @Failure 400 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/measuretypes [post]
func (h HttpServer) CreateMeasureType(c echo.Context) error {
	item := command.CreateMeasureType{}

	if err := c.Bind(&item); err != nil {
		panic(err)
	}

	if err := c.Validate(item); err != nil {
		fmt.Println(reflect.TypeOf(err))
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors[0])
		// fmt.Println(Simple(validationErrors))
		panic(errors.NewValidationError(Simple(validationErrors)))
		//return errors.NewBadRequestError(err.Error())
	}

	id, err := h.app.Commands.CreateMeasureType.Handle(c.Request().Context(), item)

	if err != nil {
		// return c.JSON(http.StatusBadRequest, err.Error())
		panic(err)
	}

	// create a response
	//return success response
	c.Response().Header().Set("location", c.Request().URL.String()+"/"+id)

	return c.JSON(http.StatusCreated, id)
}

func Simple(verr validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}
