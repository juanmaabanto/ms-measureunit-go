package command

import (
	"context"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteMeasureType struct {
	Id string
}

type DeleteMeasureTypeHandler struct {
	repo measuretype.Repository
}

func NewDeleteMeasureTypeHandler(repo measuretype.Repository) DeleteMeasureTypeHandler {
	if repo == nil {
		panic("nil repo measuretype")
	}

	return DeleteMeasureTypeHandler{repo: repo}
}

func (h DeleteMeasureTypeHandler) Handle(ctx context.Context, command DeleteMeasureType) error {
	existent := measuretype.MeasureType{}

	err := h.repo.FindById(ctx, command.Id, &existent)

	if err != nil {
		return err
	}

	if existent.Id == primitive.NilObjectID {
		return errors.NewNotFoundError("measureType")
	}

	count, err := h.repo.DeleteById(ctx, command.Id)

	if count == 0 && err == nil {
		return errors.NewBadRequestError("Failed to delete resource")
	}

	return err
}
