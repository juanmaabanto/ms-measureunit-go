package ports

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/internal/app"
	"github.com/sofisoft-tech/ms-measureunit/internal/app/command"
	"github.com/sofisoft-tech/ms-measureunit/internal/app/query"
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

// CreateMeasureType godoc
// @Summary Create a new type of measure.
// @Tags MeasureTypes
// @Accept json
// @Produce json
// @Param command body command.CreateMeasureType true "Object to be created."
// @Success 201 {string} string "Id of the created object"
// @Failure 400 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/measuretypes [post]
func (h HttpServer) AddMeasureType(c echo.Context) error {
	item := command.CreateMeasureType{}

	if err := c.Bind(&item); err != nil {
		panic(err)
	}

	if err := c.Validate(item); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		panic(errors.NewValidationError(Simple(validationErrors)))
	}

	id, err := h.app.Commands.CreateMeasureType.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	c.Response().Header().Set("location", c.Request().URL.String()+"/"+id)

	return c.JSON(http.StatusCreated, id)
}

// GetMeasureType godoc
// @Summary Get a measure type by Id.
// @Tags MeasureTypes
// @Accept json
// @Produce json
// @Param id path string  true  "MeasureType Id"
// @Success 200 {object} response.MeasureTypeResponse
// @Failure 400 {object} responses.ErrorResponse
// @Failure 404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/measuretypes/{id} [get]
func (h HttpServer) GetMeasureType(c echo.Context) error {
	item := query.GetMeasureTypeById{Id: c.Param("id")}

	result, err := h.app.Queries.GetMeasureTypeById.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, result)
}

// UpdateMeasureType godoc
// @Summary Update a type of measure.
// @Tags MeasureTypes
// @Accept json
// @Produce json
// @Param command body command.UpdateMeasureType true "Object to be modified."
// @Success 204
// @Failure 400 {object} responses.ErrorResponse
// @Failure 422 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Router /api/v1/measuretypes [patch]
func (h HttpServer) UpdateMeasureType(c echo.Context) error {
	item := command.UpdateMeasureType{}

	if err := c.Bind(&item); err != nil {
		panic(err)
	}

	if err := c.Validate(item); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		panic(errors.NewValidationError(Simple(validationErrors)))
	}

	err := h.app.Commands.UpdateMeasureType.Handle(c.Request().Context(), item)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusNoContent, nil)
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
