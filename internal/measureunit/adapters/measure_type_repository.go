package adapters

import (
	"context"

	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeasureTypeRepository struct {
	seedwork.BaseRepository
}

func NewMeasureTypeRepository(document measuretype.MeasureType) MeasureTypeRepository {
	repository := MeasureTypeRepository{
		BaseRepository: *seedwork.NewBaseRepository(&document),
	}

	return repository
}

func (repo MeasureTypeRepository) FindById(ctx context.Context, id string, receiver interface{}) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	coll := repo.BaseRepository.GetCollection()
	result := coll.FindOne(ctx, bson.D{{Key: "_id", Value: objID}})

	result.Decode(receiver)

	return nil
}
