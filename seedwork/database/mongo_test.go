package database

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockClientOpts struct {
	mock.Mock
}

func (mock *MockClientOpts) Client() *options.ClientOptions {
	args := mock.Called()
	result := args.Get(0)

	return result.(*options.ClientOptions)

}

func TestNewMongoConnection(t *testing.T) {
	// Arrange
	mockClient := new(MockClientOpts)

	mockClient.On("Client")

	// Act

	// Assert
}
