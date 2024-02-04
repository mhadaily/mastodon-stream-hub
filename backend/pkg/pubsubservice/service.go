package pubsubservice

import (
	"context"
	"log"
	"net"

	api "github.com/mhadaily/mastodon-stream-hub/pkg/api"
	"github.com/mhadaily/mastodon-stream-hub/pkg/config"
	"github.com/mhadaily/mastodon-stream-hub/pkg/pubsub"

	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedGenericPubSubServiceServer
	publisher pubsub.Publisher
}

func NewServer(publisher pubsub.Publisher) *server {
	return &server{publisher: publisher}
}

// Publish publishes a message to the PubSub
//
// how to use this function:
// 1. Create a new PubSubServiceClient
// 2. Call Publish to publish a message
// 3. Check the error
// 4. Close the stream
func (s *server) Publish(stream api.GenericPubSubService_PublishServer) error {
	for {
		log.Printf("Publishing message")
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		pubsubMsg := pubsub.Message{
			Key:   msg.Key,
			Value: msg.Value,
		}
		if err := s.publisher.Publish(context.Background(), pubsubMsg); err != nil {
			log.Printf("Failed to publish message: %v", err)
		}
	}
}

// StartServer starts the PubSubService server
// how to use
// 1. Create a new PubSubConfig
// 2. Create a new Publisher
// 3. Call StartServer to start the server
func StartServer(cfg *config.PubSubConfig, publisher pubsub.Publisher) {
	address := ":" + cfg.ServicePort
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterGenericPubSubServiceServer(s, NewServer(publisher))

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
