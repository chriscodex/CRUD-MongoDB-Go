package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// usr      = "christian"
	// pwd      = "123"
	// host     = "localhost"
	// port     = 27017
	database = "mongodb-go"
)

func GetCollection(collection string) mongo.Collection {
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)

	uri2 := "mongodb+srv://ChrisS:GprdIIKbkWpQ6mUe@clusterusers.jcbmw.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri2))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return *client.Database(database).Collection(collection)
}
