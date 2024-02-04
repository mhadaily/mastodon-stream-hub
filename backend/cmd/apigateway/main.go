package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/mhadaily/mastodon-stream-hub/pkg/config"
	"github.com/mhadaily/mastodon-stream-hub/pkg/kafkasubscriber"
	"github.com/mhadaily/mastodon-stream-hub/pkg/util"
	websocketmanager "github.com/mhadaily/mastodon-stream-hub/pkg/websocket"
)

func main() {
	cfg := config.GetPubSubConfig()

	subscriber, err := kafkasubscriber.NewKafkaSubscriber(
		cfg.KafkaBrokers,
		"mastodon-group",
		[]string{cfg.Topic},
	)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka subscriber: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup WebSocket
	http.HandleFunc("/ws", websocketmanager.HandleConnections)
	go websocketmanager.HandleMessages()

	go func() {
		log.Printf("WebSocket server started. Listening for messages...")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start WebSocket server: %v", err)
		}
	}()

	// Setup Kafka consumer
	go func() {
		log.Print("Starting Kafka consumer...")
		handler := func(msg *kafka.Message) {
			publicPost, err := util.DeserializePublicPost(msg.Value)
			if err != nil {
				log.Printf("Failed to deserialize public post: %v", err)
				return
			}

			jsonData, err := json.Marshal(publicPost)
			if err != nil {
				log.Printf("Failed to serialize public post to JSON: %v", err)
				return
			}

			jsonString := string(jsonData)
			// make sure we send json string to all connected clients via websocket using array of bytes
			websocketmanager.BroadcastMessage([]byte(jsonString))
		}
		subscriber.Listen(ctx, handler)
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down...")
	cancel()
}
