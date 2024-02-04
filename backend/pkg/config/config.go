package config

import (
	"os"
	"strings"
)

type GrpcConfig struct {
	ServiceAddr string
}

type MastodonConfig struct {
	Server       string
	ClientID     string
	ClientSecret string
	AccessToken  string
}

type PubSubConfig struct {
	PublisherType string
	KafkaBrokers  string
	Topic         string
	ServicePort   string
	GcpProjectId  string
}

type MongoDBConfig struct {
	URI            string
	DatabaseName   string
	CollectionName string
}

type ApiGatewayConfig struct {
	Port string
}

func GetApiGatewayConfig() *ApiGatewayConfig {
	return &ApiGatewayConfig{
		Port: strings.TrimSpace(os.Getenv("API_GATEWAY_SERVER_PORT")),
	}
}

func GetMongoDBConfig() *MongoDBConfig {
	return &MongoDBConfig{
		URI:            strings.TrimSpace(os.Getenv("MONGODB_URI")),
		DatabaseName:   strings.TrimSpace(os.Getenv("MONGODB_DATABASE")),
		CollectionName: strings.TrimSpace(os.Getenv("MONGODB_COLLECTION")),
	}
}

func GetPubSubConfig() *PubSubConfig {
	return &PubSubConfig{
		PublisherType: strings.TrimSpace(os.Getenv("PUBLISHER_TYPE")),
		KafkaBrokers:  strings.TrimSpace(os.Getenv("KAFKA_BROKERS")),
		Topic:         strings.TrimSpace(os.Getenv("PUBSUB_TOPIC")),
		ServicePort:   strings.TrimSpace(os.Getenv("SERVICE_PORT")),
		GcpProjectId:  strings.TrimSpace(os.Getenv("GCP_PROJECT_ID")),
	}
}

func GetMastodonConfig() *MastodonConfig {
	return &MastodonConfig{
		Server:       strings.TrimSpace(os.Getenv("MASTODON_SERVER")),
		ClientID:     strings.TrimSpace(os.Getenv("MASTODON_CLIENT_ID")),
		ClientSecret: strings.TrimSpace(os.Getenv("MASTODON_CLIENT_SECRET")),
		AccessToken:  strings.TrimSpace(os.Getenv("MASTODON_ACCESS_TOKEN")),
	}
}

func GetGrpcConfig() *GrpcConfig {
	return &GrpcConfig{
		ServiceAddr: strings.TrimSpace(os.Getenv("GRPC_SERVICE_ADDR")),
	}
}
