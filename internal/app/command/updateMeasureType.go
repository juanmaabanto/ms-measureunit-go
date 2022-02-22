package command

import (
	"context"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMeasureType struct {
	Id   string `json:"id" validate:"required,max=24,min=24"`
	Name string `json:"name" validate:"required,max=10"`
}

type UpdateMeasureTypeHandler struct {
	repo measuretype.Repository
}

func NewUpdateMeasureTypeHandler(repo measuretype.Repository) UpdateMeasureTypeHandler {
	if repo == nil {
		panic("nil repo measuretype")
	}

	return UpdateMeasureTypeHandler{repo: repo}
}

func (h UpdateMeasureTypeHandler) Handle(ctx context.Context, command UpdateMeasureType) error {
	existent := measuretype.MeasureType{}

	err := h.repo.FindById(ctx, command.Id, &existent)

	if err != nil {
		return err
	}

	if existent.Id == primitive.NilObjectID {
		return errors.NewNotFoundError("measureType")
	}

	objID, err := primitive.ObjectIDFromHex(command.Id)

	if err != nil {
		return err
	}

	count, err := h.repo.Count(ctx, bson.M{"name": primitive.Regex{
		Pattern: "^" + command.Name + "$",
		Options: "i",
	}, "_id": bson.M{"$ne": objID}})

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.NewBadRequestError("An element with the same name already exists.")
	}

	existent.Name = command.Name

	err = h.repo.UpdateOne(ctx, existent.Id.Hex(), existent)

	if err != nil {
		return err
	}

	return nil
}
