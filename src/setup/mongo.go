package setup

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoInst *Mongo

type Mongo struct {
	Client *mongo.Client
}

func NewMongo() *Mongo {
	if mongoInst != nil {
		return mongoInst
	}

	InitMongoClient()
	return mongoInst
}

func InitMongoClient() {
	if mongoInst != nil {
		return
	}

	uri := fmt.Sprintf("mongodb://%v:%v", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	opt := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	mongoInst = &Mongo{
		Client: client,
	}

	// ping to confirm connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
}
