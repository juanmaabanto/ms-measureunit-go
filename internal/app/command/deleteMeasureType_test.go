package command

import (
	"context"
	errorsN "errors"
	"testing"

	"github.com/sofisoft-tech/ms-measureunit/internal/domain/measuretype"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_NewDeleteMeasureTypeHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic")
		}
	}()

	NewDeleteMeasureTypeHandler(nil)
}

func Test_Handle_DeleteMeasureType_Find_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := DeleteMeasureType{Id: primitive.NewObjectID().Hex()}

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*measuretype.MeasureType")).Return(errorsN.New("An error has occurred"))

	// Act
	testCommand := NewDeleteMeasureTypeHandler(mockRepo)
	err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "An error has occurred")
}

func Test_Handler_DeleteMeasureType_Not_Found(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := DeleteMeasureType{Id: primitive.NewObjectID().Hex()}

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*measuretype.MeasureType")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*measuretype.MeasureType)
		arg.Id = primitive.NilObjectID
	})

	// Act
	testCommand := NewDeleteMeasureTypeHandler(mockRepo)
	err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.IsType(t, errors.ApplicationError{}, err)
	assert.Equal(t, errors.ErrorTypeNotFound, err.(errors.ApplicationError).ErrorType())
}

func Test_Handler_DeleteMeasureType_Error_Bad_Request(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := DeleteMeasureType{Id: primitive.NewObjectID().Hex()}

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*measuretype.MeasureType")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*measuretype.MeasureType)
		arg.Id = primitive.NewObjectID()
	})
	mockRepo.On("DeleteById", ctx, item.Id).Return(int64(0), nil)

	// Act
	testCommand := NewDeleteMeasureTypeHandler(mockRepo)
	err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.IsType(t, errors.ApplicationError{}, err)
	assert.Equal(t, errors.ErrorTypeBadRequest, err.(errors.ApplicationError).ErrorType())
}

func Test_Handler_DeleteMeasureType_Error_Completed(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := DeleteMeasureType{Id: primitive.NewObjectID().Hex()}

	mockRepo.On("FindById", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("*measuretype.MeasureType")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(2).(*measuretype.MeasureType)
		arg.Id = primitive.NewObjectID()
	})
	mockRepo.On("DeleteById", ctx, item.Id).Return(int64(1), nil)

	// Act
	testCommand := NewDeleteMeasureTypeHandler(mockRepo)
	err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
}
