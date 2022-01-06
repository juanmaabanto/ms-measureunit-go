package service

import (
	"context"

	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/adapters"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app/command"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/domain/measuretype"
)

func NewApplication(ctx context.Context) app.Application {
	document := new(measuretype.MeasureType)

	measureTypeRepository := adapters.NewMeasureTypeRepository(*document)

	return app.Application{
		Commands: app.Commands{
			CreateMeasureType: command.NewCreateMeasureTypeHandler(measureTypeRepository),
		},
	}
}
