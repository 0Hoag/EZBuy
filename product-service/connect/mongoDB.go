package connect

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	db   *mongo.Client
	once sync.Once
)

func InitDatabase() (*mongo.Client, error) {
	var err error
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		db, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		err = db.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("Failed to ping MongoDB: %v", err)
		}

		log.Println("Successfully connected to MongoDB")

		database := db.Database("product-service")
		_ = database.Collection("category")
		_ = database.Collection("product")
	})

	return db, err
}

func getDB() *mongo.Client {
	return db
}
