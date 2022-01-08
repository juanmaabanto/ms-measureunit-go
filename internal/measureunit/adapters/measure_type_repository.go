package adapters

import (
	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/database"
)

type MeasureTypeRepository struct {
	seedwork.BaseRepository
}

func NewMeasureTypeRepository(connection database.MongoConnection, document measuretype.MeasureType) MeasureTypeRepository {
	repository := MeasureTypeRepository{
		BaseRepository: *seedwork.NewBaseRepository(connection, &document),
	}

	return repository
}
