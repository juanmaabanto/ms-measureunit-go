package service

import (
	"context"
	"os"

	"github.com/sofisoft-tech/ms-measureunit/internal/app"
	"github.com/sofisoft-tech/ms-measureunit/internal/app/command"
	"github.com/sofisoft-tech/ms-measureunit/internal/app/query"
	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/internal/infrastructure"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/database"
)

func NewApplication(ctx context.Context) app.Application {
	conn := database.NewMongoConnection(ctx, os.Getenv("MONGODB_NAME"), os.Getenv("MONGODB_URI"))
	document := new(measuretype.MeasureType)

	measureTypeRepository := infrastructure.NewMeasureTypeRepository(conn, *document)

	return app.Application{
		Commands: app.Commands{
			CreateMeasureType: command.NewCreateMeasureTypeHandler(measureTypeRepository),
			UpdateMeasureType: command.NewUpdateMeasureTypeHandler(measureTypeRepository),
		},
		Queries: app.Queries{
			GetMeasureTypeById: query.NewGetMeasureTypeByIdHandler(measureTypeRepository),
		},
	}
}
