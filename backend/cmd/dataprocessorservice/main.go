package main

import (
	"github.com/mhadaily/mastodon-stream-hub/pkg/config"
	dataprocessor "github.com/mhadaily/mastodon-stream-hub/pkg/dataprocessor"
)

func main() {
	port := config.GetPubSubConfig().ServicePort
	dataprocessor.StartServer(port)

	// This is the main entry point for the DataProcessor service.
	// potentially we can add more logic here.
	// like connecting with other services
}
