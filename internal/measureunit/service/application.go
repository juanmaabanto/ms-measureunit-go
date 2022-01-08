package service

import (
	"context"
	"os"

	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/adapters"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app/command"
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/database"
)

func NewApplication(ctx context.Context) app.Application {
	conn := database.NewMongoConnection(ctx, os.Getenv("MONGODB_NAME"), os.Getenv("MONGODB_URI"))
	document := new(measuretype.MeasureType)

	measureTypeRepository := adapters.NewMeasureTypeRepository(conn, *document)

	return app.Application{
		Commands: app.Commands{
			CreateMeasureType: command.NewCreateMeasureTypeHandler(measureTypeRepository),
		},
	}
}
