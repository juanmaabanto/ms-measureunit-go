package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/responses"
)

type LoggerConfig struct {
	LoggerErrorFunc func(message string, trace string, username string) string
}

var DefaultLoggerConfig = LoggerConfig{
	LoggerErrorFunc: func(message string, trace string, username string) string {
		fmt.Println(message)

		return "Id autogenerado por defaultLoggerConfig"
	},
}

func Logger() echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}

func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					customErr, ok := err.(errors.ApplicationError)

					if ok && customErr.ErrorType() == errors.ErrorTypeBadRequest {
						response := responses.ErrorResponse{
							Message: customErr.Error(),
							Status:  400,
							Title:   customErr.Title(),
						}

						c.JSON(http.StatusBadRequest, response)
					} else {
						errorId := config.LoggerErrorFunc(err.Error(), "trace", "test")

						response := responses.ErrorResponse{
							ErrorId: errorId,
							Message: err.Error(),
							Status:  500,
						}
						c.JSON(http.StatusInternalServerError, response)
					}
				}
			}()

			return next(c)
		}
	}
}
