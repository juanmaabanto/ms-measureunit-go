package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb+srv://root:A123456a@cluster0.sf8k2.mongodb.net/MS-MeasureUnits?retryWrites=true&w=majority"

var (
	Database *mongo.Database
)

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged to mongo.")

	Database = client.Database("MS-MeasureUnits")
}
