package main

import (
	"context"
	"log"

	api "github.com/mhadaily/mastodon-stream-hub/pkg/api"
	"github.com/mhadaily/mastodon-stream-hub/pkg/config"
	"github.com/mhadaily/mastodon-stream-hub/pkg/mastodonclient"
	"github.com/mhadaily/mastodon-stream-hub/pkg/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Printf("Mastodon Stream Service")
	grpcConfig := config.GetGrpcConfig()

	// Set up the gRPC connection to the PubSubService
	conn, err := grpc.Dial(
		grpcConfig.ServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// Create a new PubSubServiceClient
	pubSubClient := api.NewGenericPubSubServiceClient(conn)
	stream, err := pubSubClient.Publish(context.Background())
	if err != nil {
		log.Fatalf("Error opening publish stream: %v", err)
	}

	mastodonConfig := config.GetMastodonConfig()
	client := mastodonclient.NewClient(mastodonConfig)

	// Define a handler function for new public posts
	//
	// The handler function is called for each new public post
	handler := func(post *api.PublicPost) {
		msg, err := util.CreateGenericMessage(post)
		if err != nil {
			log.Fatalf("Failed to create generic message: %v", err)
			return
		}
		log.Printf("message: %v", post.Id)
		if err := stream.Send(msg); err != nil {
			log.Fatalf("Failed to send post over stream: %v", err)
		}
	}

	// Stream public posts from the Mastodon server
	// and publish them to the PubSubService
	//
	// check *api.PublicPost in backend/pkg/api/mastodonstream.pb.go
	// which is passed to the handler function
	err = client.StreamPublicPosts(context.Background(), handler)
	if err != nil {
		log.Fatalf("Error streaming public posts: %v", err)
	}
}
