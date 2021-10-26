package clients

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var client *mongo.Client

func NewMongoClient(ctx context.Context) *mongo.Client {
/*	if client != nil {
		return client
	}

	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)*/

	return nil
}

func CloseMongoClient(ctx context.Context) {
	if client == nil {
		return
	}

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB clients
	fmt.Println("Connection to MongoDB closed.")
}