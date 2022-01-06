package app

import "github.com/sofisoft-tech/ms-measureunit/internal/measureunit/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	CreateMeasureType command.CreateMeasureTypeHandler
}
