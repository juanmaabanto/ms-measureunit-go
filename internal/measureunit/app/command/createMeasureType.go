package command

import (
	"context"
	"errors"
	"time"

	"github.com/sofisoft-tech/ms-measureunit/internal/measureunit/domain/measuretype"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	count, err := h.repo.Count(ctx, bson.M{"name": primitive.Regex{
		Pattern: "^" + command.Name + "$",
		Options: "i",
	}})

	if count > 0 {
		return "", errors.New("Un error para el logger")
		// return "", errors.NewBadRequestError("Ya existe un elemento con el mismo nombre") //.New("Ya existe un elemento con el mismo nombre")
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
