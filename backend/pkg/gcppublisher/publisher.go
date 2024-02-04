package gcppublisher

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	api "github.com/mhadaily/mastodon-stream-hub/pkg/pubsub"
)

type GCPPublisher struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

// NewGCPPublisher creates a new GCP Pub/Sub publisher.
//
// how to use
// publisher := gcppublisher.NewGCPPublisher(
//
//	cfg.GcpProjectId,
//	cfg.Topic,
//
// )
func NewGCPPublisher(projectID string, topicID string) *GCPPublisher {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	topic := client.Topic(topicID)

	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Failed to verify Pub/Sub topic exists: %v", err)
	}
	if !exists {
		log.Fatalf("Topic %s does not exist", topicID)
	}

	return &GCPPublisher{
		client: client,
		topic:  topic,
	}
}

// Publish sends a message to the GCP Pub/Sub topic.
// The message key is used as an attribute in the Pub/Sub message.
// The message value is the payload of the Pub/Sub message.
// how to use
// publisher := gcppublisher.NewGCPPublisher(
//
//	cfg.GcpProjectId,
//	cfg.Topic,
//
// )
func (g *GCPPublisher) Publish(ctx context.Context, msg api.Message) error {
	// The Pub/Sub library manages batching and sending messages.
	result := g.topic.Publish(ctx, &pubsub.Message{
		Data:       msg.Value,
		Attributes: map[string]string{"key": string(msg.Key)},
	})

	// Get blocks until the result is available.
	_, err := result.Get(ctx)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
	}
	return err
}

// Close cleans up any resources held by the publisher.
// In this case, it closes the Pub/Sub client.
func (g *GCPPublisher) Close() error {
	// There's no explicit close method for the client or topic in the Pub/Sub SDK,
	// but you could manage any cleanup tasks here if necessary.
	return nil
}
