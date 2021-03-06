package command

import (
	"context"
	"time"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMeasureType struct {
	Name string `json:"name" validate:"required,max=10"`
}

type CreateMeasureTypeHandler struct {
	repo measuretype.Repository
}

func NewCreateMeasureTypeHandler(repo measuretype.Repository) CreateMeasureTypeHandler {
	if repo == nil {
		panic("nil repo measuretype")
	}

	return CreateMeasureTypeHandler{repo: repo}
}

func (h CreateMeasureTypeHandler) Handle(ctx context.Context, command CreateMeasureType) (string, error) {
	count, err := h.repo.Count(ctx, bson.M{"name": primitive.Regex{
		Pattern: "^" + command.Name + "$",
		Options: "i",
	}})

	if err != nil {
		return "", err
	}

	if count > 0 {
		return "", errors.NewBadRequestError("An element with the same name already exists.")
	}

	item := measuretype.MeasureType{}

	item.Name = command.Name
	item.CreatedAt = time.Now()
	item.CreatedBy = "admin"

	id, err := h.repo.InsertOne(ctx, item)

	if err != nil {
		return id, err
	}

	return id, nil
}
