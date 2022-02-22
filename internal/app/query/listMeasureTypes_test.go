package query

import (
	"context"
	"errors"
	"testing"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NewListMeasureTypesHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic")
		}
	}()

	NewListMeasureTypesHandler(nil)
}

func Test_Handle_ListMeasureTypes_Count_Err(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.D")).Return(int64(0), errors.New("An error"))

	// Act
	testQuery := NewListMeasureTypesHandler(mockRepo)
	_, _, err := testQuery.Handle(ctx, ListMeasureTypes{Name: "", Start: 0, PageSize: 50})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, "An error", err.Error())
}

func Test_Handle_ListMeasureTypes_Paginated_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.D")).Return(int64(1), nil)
	mockRepo.On("Paginated", ctx, mock.AnythingOfType("primitive.D"), mock.AnythingOfType("primitive.D"), int64(50), int64(0), mock.AnythingOfType("*[]measuretype.MeasureType")).Return(errors.New("An error"))

	// Act
	testQuery := NewListMeasureTypesHandler(mockRepo)
	_, _, err := testQuery.Handle(ctx, ListMeasureTypes{Name: "", Start: 0, PageSize: 50})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, "An error", err.Error())
}

func Test_Handle_ListMeasureTypes_Ok(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.D")).Return(int64(1), nil)
	mockRepo.On("Paginated", ctx, mock.AnythingOfType("primitive.D"), mock.AnythingOfType("primitive.D"), int64(50), int64(0), mock.AnythingOfType("*[]measuretype.MeasureType")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(5).(*[]measuretype.MeasureType)

		*arg = append(*arg, measuretype.MeasureType{
			Name: "test",
		})

	})

	// Act
	testQuery := NewListMeasureTypesHandler(mockRepo)
	total, results, _ := testQuery.Handle(ctx, ListMeasureTypes{Name: "", Start: 0, PageSize: 50})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Equal(t, int64(1), total)
	assert.Equal(t, "test", results[0].Name)
}
