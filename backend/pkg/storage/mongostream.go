package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChangeStreamHandler func(data bson.M)

// WatchCollection watches for changes in a MongoDB collection and calls the handler function for each change.
// how to use it:
// storage.WatchCollection(collection, func(data bson.M) {
// })
func WatchCollection(collection *mongo.Collection, handler ChangeStreamHandler) {
	pipeline := mongo.Pipeline{bson.D{{Key: "$match", Value: bson.D{{Key: "operationType", Value: "insert"}}}}}
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	changeStream, err := collection.Watch(context.Background(), pipeline, opts)
	if err != nil {
		log.Fatalf("Failed to watch collection changes: %v", err)
	}
	defer changeStream.Close(context.Background())

	for changeStream.Next(context.Background()) {
		var data bson.M
		if err := changeStream.Decode(&data); err != nil {
			log.Fatalf("Failed to decode change stream data: %v", err)
		}
		handler(data)
	}
}
