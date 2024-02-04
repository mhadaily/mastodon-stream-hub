package main

import (
	"log"

	"github.com/mhadaily/mastodon-stream-hub/pkg/config"
	"github.com/mhadaily/mastodon-stream-hub/pkg/gcppublisher"
	"github.com/mhadaily/mastodon-stream-hub/pkg/kafkapublisher"
	"github.com/mhadaily/mastodon-stream-hub/pkg/pubsub"
	"github.com/mhadaily/mastodon-stream-hub/pkg/pubsubservice"
)

func main() {
	cfg := config.GetPubSubConfig()

	var publisher pubsub.Publisher
	switch cfg.PublisherType {
	case "KAFKA":
		brokers := cfg.KafkaBrokers
		topic := cfg.Topic
		publisher = kafkapublisher.NewKafkaPublisher(brokers, topic)
	case "GCP_PUBSUB":
		projectID := cfg.GcpProjectId
		topicID := cfg.Topic
		publisher = gcppublisher.NewGCPPublisher(projectID, topicID)
	default:
		log.Fatalf("Unsupported publisher type: %s", cfg.PublisherType)
	}

	pubsubservice.StartServer(cfg, publisher)
}
