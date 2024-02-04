package kafkapublisher

import (
	"context"
	"log"

	"github.com/mhadaily/mastodon-stream-hub/pkg/pubsub"

	kafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaPublisher struct {
	producer *kafka.Producer
	topic    string
}

// NewKafkaPublisher creates a new Kafka publisher.
//
// how to use
// publisher := kafkapublisher.NewKafkaPublisher(
//
//	cfg.KafkaBrokers,
//	cfg.Topic,
//
// )
func NewKafkaPublisher(brokers string, topic string) *KafkaPublisher {
	log.Printf("Configuring Kafka producer with bootstrap.servers: %s", brokers)
	config := &kafka.ConfigMap{"bootstrap.servers": brokers}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	return &KafkaPublisher{
		producer: producer,
		topic:    topic,
	}
}

// Publish sends a message to the Kafka topic.
// The message key is used as an attribute in the Kafka message.
// The message value is the payload of the Kafka message.
//
// how to use
// publisher := kafkapublisher.NewKafkaPublisher(
//
//	cfg.KafkaBrokers,
//	cfg.Topic,
//
// )
func (k *KafkaPublisher) Publish(ctx context.Context, msg pubsub.Message) error {
	kafkaMsg := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &k.topic,
			Partition: kafka.PartitionAny,
		},
		Key:   msg.Key,
		Value: msg.Value,
	}
	return k.producer.Produce(&kafkaMsg, nil)
}

func (k *KafkaPublisher) Close() error {
	k.producer.Close()
	return nil // Confluent Kafka producer close method does not return an error
}
