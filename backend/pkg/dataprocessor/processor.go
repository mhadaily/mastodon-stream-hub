package dataprocessor

import (
	"context"
	"log"
	"net"

	api "github.com/mhadaily/mastodon-stream-hub/pkg/api"
	"github.com/mhadaily/mastodon-stream-hub/pkg/util"
	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedDataProcessorServer
}

func NewServer() *server {
	return &server{}
}

// ProcessMessage processes a single message from Kafka
//
// This is just a mimic of a real-world processing step
// Ideally I should have used a real processing step here like Apache Flink, Apache Beam, or Apache Spark
// But dude to lack of time, I had to skip this part.
func ProcessData(
	ctx context.Context,
	request *api.PubSubGenericMessage,
) (*api.PublicPost, error) {
	// Deserialize the message to a PublicPost
	// This is just a mimic of a real-world processing step
	post, err := util.DeserializePublicPost(request.Value)
	if err != nil {
		log.Fatalf("Failed to deserialize message: %v", err)
	}
	log.Printf("Processing Post ID: %s", post.Id)

	return post, nil
}

func StartServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterDataProcessorServer(s, NewServer())
	log.Printf("DataProcessor server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
