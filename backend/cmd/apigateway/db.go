package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mhadaily/mastodon-stream-hub/pkg/config"
	"github.com/mhadaily/mastodon-stream-hub/pkg/storage"
	websocketmanager "github.com/mhadaily/mastodon-stream-hub/pkg/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// If Replication is enabled, MongoDB can be used to store data in a distributed manner.
// In this case it can also help for high availability and fault tolerance.
// as well as realtime data checking.
func withMonoDB() {
	mongoConfig := config.GetMongoDBConfig()
	apiGatewayConfig := config.GetApiGatewayConfig()

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(mongoConfig.URI),
	)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(mongoConfig.DatabaseName).Collection(mongoConfig.CollectionName)

	// MongoDB can provide a change stream for a collection.
	// This allows us to watch for changes in the collection and react to them.
	go storage.WatchCollection(collection, func(data bson.M) {
		websocketmanager.BroadcastMessage([]byte(fmt.Sprintf("%v", data)))
	})

	// Then we can start the WebSocket server.
	// The WebSocket server listens for incoming connections and broadcasts messages to all connected clients.
	http.HandleFunc("/ws", websocketmanager.HandleConnections)
	go websocketmanager.HandleMessages()

	log.Printf("WebSocket server started on %v", apiGatewayConfig.Port)
	// start server
	err = http.ListenAndServe(apiGatewayConfig.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
