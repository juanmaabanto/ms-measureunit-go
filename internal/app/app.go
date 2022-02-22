package app

import (
	"github.com/sofisoft-tech/ms-measureunit/internal/app/command"
	"github.com/sofisoft-tech/ms-measureunit/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateMeasureType command.CreateMeasureTypeHandler
}

type Queries struct {
	GetMeasureTypeById query.GetMeasureTypeByIdHandler
}
