package storage

import (
	"context"
	"log"

	api "github.com/mhadaily/mastodon-stream-hub/pkg/api"
	"github.com/mhadaily/mastodon-stream-hub/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStorage struct {
	collection *mongo.Collection
}

// NewMongoDBStorage initializes a new instance of MongoDBStorage with the given configuration.
// it accepts Config struct and returns a pointer to MongoDBStorage and an error.
// how to use it:
// cfg := config.GetMongoDBConfig()
// storage, err := storage.NewMongoDBStorage(cfg)
//
//	if err != nil {
//		log.Fatalf("Failed to initialize MongoDB storage: %v", err)
//	}
func NewMongoDBStorage(cfg *config.MongoDBConfig) (*MongoDBStorage, error) {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(cfg.URI),
	)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	collection := client.Database(cfg.DatabaseName).Collection(cfg.CollectionName)
	return &MongoDBStorage{collection: collection}, nil
}

// StorePost saves a processed post to MongoDB.
// InsertOne is a method of the Collection type in the mongo-go-driver package.
// how to use it:
// err := storage.StorePost(ctx, post)
//
//	if err != nil {
//		log.Fatalf("Failed to store post: %v", err)
//	}
func (s *MongoDBStorage) StorePost(ctx context.Context, post *api.PublicPost) error {
	_, err := s.collection.InsertOne(ctx, post)
	if err != nil {
		log.Printf("Failed to insert post: %v", err)
		return err
	}
	log.Println("Post stored successfully")
	return nil
}
