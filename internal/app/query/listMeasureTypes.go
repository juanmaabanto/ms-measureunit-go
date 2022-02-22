package query

import (
	"context"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/internal/ports/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListMeasureTypes struct {
	Name     string
	Start    int64
	PageSize int64
}

type ListMeasureTypesHandler struct {
	repo measuretype.Repository
}

func NewListMeasureTypesHandler(repo measuretype.Repository) ListMeasureTypesHandler {
	if repo == nil {
		panic("nil repo")
	}

	return ListMeasureTypesHandler{repo}
}

func (h ListMeasureTypesHandler) Handle(ctx context.Context, query ListMeasureTypes) (int64, []response.MeasureTypeResponse, error) {
	var items []measuretype.MeasureType
	results := []response.MeasureTypeResponse{}

	filter := bson.D{
		{"$or",
			bson.A{
				bson.D{{"name", primitive.Regex{
					Pattern: query.Name,
					Options: "i",
				}}},
			}},
	}

	total, err := h.repo.Count(ctx, filter)

	if err != nil {
		return 0, results, err
	}

	err = h.repo.Paginated(ctx, filter, bson.D{}, query.PageSize, query.Start, &items)

	if err != nil {
		return 0, results, err
	}

	for _, element := range items {
		results = append(results, response.MeasureTypeResponse{
			Id:   element.Id.Hex(),
			Name: element.Name,
		})
	}

	return total, results, err
}
