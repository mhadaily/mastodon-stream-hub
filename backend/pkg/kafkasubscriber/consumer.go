package kafkasubscriber

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaSubscriber struct {
	consumer *kafka.Consumer
}

// NewKafkaSubscriber initializes a new Kafka consumer.
//
// It accepts a list of brokers, a group ID, and a list of topics to subscribe to.
// It returns a pointer to KafkaSubscriber and an error.
// how to use:
// subscriber, err := kafkasubscriber.NewKafkaSubscriber(
//
//	cfg.KafkaBrokers,
//	"mastodon-group",
//	[]string{cfg.Topic},
//
// )
//
//	if err != nil {
//		log.Fatalf("Failed to initialize Kafka subscriber: %v", err)
//	}
func NewKafkaSubscriber(brokers string, groupID string, topics []string) (*KafkaSubscriber, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %v", err)
	}

	log.Printf("Subscribing to topics: %v", topics)
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topics: %v", err)
	}

	return &KafkaSubscriber{
		consumer: consumer,
	}, nil
}

// Listen listens for messages on the subscribed topics.
// The provided handler function is called for each message.
// how to use it:
// subscriber.Listen(ctx, handler)
func (s *KafkaSubscriber) Listen(ctx context.Context, handler func(*kafka.Message)) {
	run := true
	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			// The call will block for at most `timeout` waiting for
			// a new message or error. `timeout` may be set to -1 for
			// indefinite wait.
			msg, err := s.consumer.ReadMessage(-1)
			if err == nil {
				handler(msg)
			} else if e, ok := err.(kafka.Error); ok && !e.IsFatal() {
				log.Printf("Consumer non-fatal error: %v\n", err)
			} else {
				log.Printf("Consumer error: %v\n", err)
				run = false
			}
		}
	}

	log.Println("Closing Kafka consumer...")
	if err := s.consumer.Close(); err != nil {
		log.Printf("Failed to close Kafka consumer: %v", err)
	}
}
