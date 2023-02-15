package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect")
		return nil
	}

	fmt.Println("Successfully connected to mongodb")
	return client

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
  var collection *mongo.Collection = client.Database("dcommerce").Collection(collectionName)
  return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("dcommerce").Collection(collectionName)
	return productcollection
}
