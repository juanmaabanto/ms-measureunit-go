package command

import (
	"context"
	errorsN "errors"
	"testing"

	"github.com/sofisoft-tech/ms-measureunit/seedwork/errors"
	"github.com/sofisoft-tech/ms-measureunit/seedwork/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_NewCreateMeasureTypeHandler(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic")
		}
	}()

	NewCreateMeasureTypeHandler(nil)
}

func Test_Handle_CreateMeasureType_Count_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := CreateMeasureType{Name: "test"}

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.M")).Return(int64(0), errorsN.New("An error has occurred"))

	// Act
	testCommand := NewCreateMeasureTypeHandler(mockRepo)
	_, err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "An error has occurred")
}

func Test_Handle_CreateMeasureType_Count_Is_Greater_Than_Zero(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := CreateMeasureType{Name: "test"}

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.M")).Return(int64(1), nil)

	// Act
	testCommand := NewCreateMeasureTypeHandler(mockRepo)
	_, err := testCommand.Handle(ctx, item)

	// Assert
	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.IsType(t, errors.ApplicationError{}, err)
	assert.Equal(t, errors.ErrorTypeBadRequest, err.(errors.ApplicationError).ErrorType())
}

func Test_Handle_CreateMeasureType_Insert_Completed(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := CreateMeasureType{Name: "test"}
	expected := primitive.NewObjectID().Hex()

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.M")).Return(int64(0), nil)
	mockRepo.On("InsertOne", ctx, mock.AnythingOfType("measuretype.MeasureType")).Return(expected, nil)

	// Act
	testCommand := NewCreateMeasureTypeHandler(mockRepo)
	result, err := testCommand.Handle(ctx, item)

	// Assert

	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func Test_Handle_CreateMeasureType_Insert_Error(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockRepository)
	ctx := context.Background()
	item := CreateMeasureType{Name: "test"}

	mockRepo.On("Count", ctx, mock.AnythingOfType("primitive.M")).Return(int64(0), nil)
	mockRepo.On("InsertOne", ctx, mock.AnythingOfType("measuretype.MeasureType")).Return("", errorsN.New("An error has occurred"))

	// Act
	testCommand := NewCreateMeasureTypeHandler(mockRepo)
	_, err := testCommand.Handle(ctx, item)

	// Assert

	mockRepo.AssertExpectations(t)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "An error has occurred")
}
