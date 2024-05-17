package repository

import (
	"context"
	"crud_operation/configs"
	"fmt"
	"log"
	"net/url"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	StorageInstance *mongo.Client
	once            sync.Once
)

func Init(cfg configs.ApiConfig) error {
	var err error

	once.Do(func() {
		log.Println("Connecting to MongoDB...")

		// Encode the username and password to be URL-safe
		encodedUsername := url.QueryEscape(cfg.MongoUsername)
		encodedPassword := url.QueryEscape(cfg.MongoPassword)

		mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", encodedUsername, encodedPassword, cfg.MongoDBURI)

		// Set client options
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}

		// Check the connection
		err = client.Ping(context.Background(), nil)
		if err != nil {
			panic(err)
		}

		// Send a ping to confirm a successful connection
		if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
			panic(err)
		}

		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

		log.Println("Connected to MongoDB!")
		StorageInstance = client
	})

	return err
}

// getting database collections
func GetCollection(cfg configs.ApiConfig, client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(cfg.MongoDBName).Collection(collectionName)
	return collection
}
