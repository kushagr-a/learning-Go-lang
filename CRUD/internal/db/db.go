package db

import (
	"context"
	"fmt"
	"note-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(cfg config.Config) (*mongo.Client, *mongo.Database, error) {

	// prevent app freeze on startup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect options
	clientOps := options.Client().ApplyURI(cfg.MongoURI)

	// connect to mongo
	client, err := mongo.Connect(ctx, clientOps)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// ping db
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	database := client.Database(cfg.MongoDB)

	fmt.Println("Connected to MongoDB")

	return client, database, nil
}

func DisconnectDB(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}
