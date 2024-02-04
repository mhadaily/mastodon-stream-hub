package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	api "github.com/mhadaily/mastodon-stream-hub/pkg/api"
	"github.com/mhadaily/mastodon-stream-hub/pkg/config"
	"github.com/mhadaily/mastodon-stream-hub/pkg/dataprocessor"
	"github.com/mhadaily/mastodon-stream-hub/pkg/kafkasubscriber"
	"github.com/mhadaily/mastodon-stream-hub/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.GetPubSubConfig()
	grpcConfig := config.GetGrpcConfig()
	mongoConfig := config.GetMongoDBConfig()

	mongoStorage, err := storage.NewMongoDBStorage(mongoConfig)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB storage: %v", err)
	}

	log.Print("Starting Kafka consumer...", cfg.KafkaBrokers, cfg.Topic)
	topics := []string{cfg.Topic}
	subscriber, err := kafkasubscriber.NewKafkaSubscriber(
		cfg.KafkaBrokers,
		"mastodon-group",
		topics,
	)
	if err != nil {
		log.Fatalf("Failed to initialize Kafka subscriber: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		log.Println("Shutting down...")
		cancel()
	}()

	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(
		grpcConfig.ServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// processDataGrpcClient := api.NewDataProcessorClient(conn)

	handler := func(msg *kafka.Message) {
		log.Printf("Received message: %s", msg)
		// Process message using gRPC client
		// publicPost, err := processDataGrpcClient.ProcessData(
		// 	context.Background(),
		// 	&api.PubSubGenericMessage{Value: msg.Value, Key: msg.Key},
		// )

		publicPost, err := dataprocessor.ProcessData(
			context.Background(),
			&api.PubSubGenericMessage{Value: msg.Value, Key: msg.Key},
		)
		if err != nil {
			log.Printf("Failed to process data: %v", err)
			return
		}

		// Store processed data in MongoDB
		log.Printf("Storing processed data %v", publicPost.Id)
		if err := mongoStorage.StorePost(context.Background(), publicPost); err != nil {
			log.Printf("Failed to store processed data: %v", err)
		}
	}

	log.Print("Listening for messages...")
	// Start listening for messages
	subscriber.Listen(ctx, handler)
}
