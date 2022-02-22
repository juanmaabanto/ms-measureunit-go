package query

import (
	"context"
	errorsN "errors"
	"testing"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/internal/ports/response"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_NewGetMeasureTypeByIdHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic")
		}
	}()

	NewGetMeasureTypeByIdHandler(nil)
}

func Test_Handle_GetMeasureTypeById_Found(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	expected := response.MeasureTypeResponse{Id: "6209a7fb0ceab9da565c546d", Name: "test"}

	mockRepo.On("FindById", ctx, expected.Id, mock.AnythingOfType("*measuretype.MeasureType")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*measuretype.MeasureType)
		arg.Id, _ = primitive.ObjectIDFromHex(expected.Id)
		arg.Name = expected.Name
	})

	// Act
	testQuery := NewGetMeasureTypeByIdHandler(mockRepo)
	result, err := testQuery.Handle(ctx, GetMeasureTypeById{Id: expected.Id})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Equal(t, expected.Id, result.Id)
	assert.Equal(t, expected.Name, result.Name)
	assert.Nil(t, err)
}

func Test_Handler_GetMeasureTypeById_Not_Found(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*measuretype.MeasureType")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*measuretype.MeasureType)
		arg.Id = primitive.NilObjectID
	})

	// Act
	testQuery := NewGetMeasureTypeByIdHandler(mockRepo)
	_, err := testQuery.Handle(ctx, GetMeasureTypeById{Id: primitive.NewObjectID().String()})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.IsType(t, errors.ApplicationError{}, err)
	assert.Equal(t, errors.ErrorTypeNotFound, err.(errors.ApplicationError).ErrorType())
}

func Test_Handler_GetMeasureTypeById_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*measuretype.MeasureType")).Return(errorsN.New("An error has occurred"))

	// Act
	testQuery := NewGetMeasureTypeByIdHandler(mockRepo)
	_, err := testQuery.Handle(ctx, GetMeasureTypeById{Id: primitive.NewObjectID().String()})

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "An error has occurred")
}
