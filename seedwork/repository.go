package seedwork

import (
	"context"

	db "github.com/sofisoft-tech/ms-measureunit/seedwork/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBaseRepository interface {
	GetById(ctx context.Context, id string, receiver interface{}) error
	InsertOne(ctx context.Context, document IDocument) (string, error)
}

type BaseRepository struct {
	collection *mongo.Collection
}

func NewBaseRepository(document IDocument) *BaseRepository {
	repository := &BaseRepository{
		collection: db.Database.Collection(document.GetCollectionName()),
	}

	return repository
}

func (repository BaseRepository) GetById(ctx context.Context, id string, receiver interface{}) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}

	coll := repository.collection
	result := coll.FindOne(ctx, bson.D{{Key: "_id", Value: objID}})

	result.Decode(receiver)

	return nil
}

func (repository BaseRepository) InsertOne(ctx context.Context, document IDocument) (string, error) {
	coll := repository.collection

	result, err := coll.InsertOne(ctx, document)

	return result.InsertedID.(primitive.ObjectID).String(), err
}

func (repository BaseRepository) GetCollection() *mongo.Collection {
	return repository.collection
}
