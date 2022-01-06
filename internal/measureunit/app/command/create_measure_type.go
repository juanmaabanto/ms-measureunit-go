package command

import (
	"context"
	"time"

	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/domain/measuretype"
)

type CreateMeasureType struct {
	Name string `json:"name"`
}

type CreateMeasureTypeHandler struct {
	repo measuretype.Repository
}

func NewCreateMeasureTypeHandler(repo measuretype.Repository) CreateMeasureTypeHandler {
	if repo == nil {
		panic("nil repo")
	}

	return CreateMeasureTypeHandler{repo: repo}
}

func (h CreateMeasureTypeHandler) Handle(ctx context.Context, command CreateMeasureType) (string, error) {
	item := new(measuretype.MeasureType)

	item.Name = command.Name
	item.CreatedAt = time.Now()
	item.CreatedBy = "admin"

	id, err := h.repo.InsertOne(ctx, item)

	if err != nil {
		return id, err
	}

	return id, nil
}
