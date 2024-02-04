package pubsub

import "context"

// Message defines the structure of a message
// that can be published to a pub/sub system
// such as Kafka or Google Cloud Pub/Sub
//
// Key: The key of the message
// Value: The value of the message
type Message struct {
	Key   []byte
	Value []byte
}

// Publisher is an interface that defines the
// methods that a pub/sub publisher should implement
//
// Publish: Publishes a message to the pub/sub system
// Close: Closes the connection to the pub/sub system
type Publisher interface {
	Publish(ctx context.Context, msg Message) error
	Close() error
}
