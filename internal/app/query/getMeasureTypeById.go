package query

import (
	"context"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/internal/ports/response"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetMeasureTypeById struct {
	Id string
}

type GetMeasureTypeByIdHandler struct {
	repo measuretype.Repository
}

func NewGetMeasureTypeByIdHandler(repo measuretype.Repository) GetMeasureTypeByIdHandler {
	if repo == nil {
		panic("nil repo")
	}

	return GetMeasureTypeByIdHandler{repo}
}

func (h GetMeasureTypeByIdHandler) Handle(ctx context.Context, query GetMeasureTypeById) (*response.MeasureTypeResponse, error) {
	receiver := measuretype.MeasureType{}

	err := h.repo.FindById(ctx, query.Id, &receiver)

	if err != nil {
		return nil, err
	}

	if receiver.Id == primitive.NilObjectID {
		return nil, errors.NewNotFoundError("measureType")
	}

	response := response.MeasureTypeResponse{
		Id:   receiver.Id.Hex(),
		Name: receiver.Name,
	}

	return &response, nil
}
