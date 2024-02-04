package mastodonclient

import (
	"context"
	"log"

	api "github.com/mhadaily/mastodon-stream-hub/pkg/api"
	"github.com/mhadaily/mastodon-stream-hub/pkg/config"

	"github.com/mattn/go-mastodon"
	"github.com/mitchellh/mapstructure"
)

type StreamHandler func(*api.PublicPost)

type Client struct {
	Mastodon *mastodon.Client
}

func NewClient(config *config.MastodonConfig) *Client {
	return &Client{
		Mastodon: mastodon.NewClient(&mastodon.Config{
			Server:       config.Server,
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			AccessToken:  config.AccessToken,
		}),
	}
}

// StreamPublicPosts streams public posts from the Mastodon server
//
// The handler function is called for each new public post
// ctx is the context for the stream
// The function returns an error if the stream fails
// how to use
// 1. Create a new MastodonConfig
// 2. Create a new Client
// 3. Call StreamPublicPosts to start streaming public posts
// 4. Pass a handler function to handle the posts
// 5. Check the error
// 6. Handle the posts
// 7. Close the stream
func (c *Client) StreamPublicPosts(
	ctx context.Context,
	handler StreamHandler,
) error {
	// Create a new WebSocket client
	//
	// While server-sent events (SSE) are also an option
	// the websocket was chosen over the REST API because it is more efficient
	wsClient := c.Mastodon.NewWSClient()

	// Connect to the streaming API
	//
	// The second parameter is set to false to receive all public posts
	// it was chose to stress the architecture of the system
	publicStream, err := wsClient.StreamingWSPublic(ctx, false)
	if err != nil {
		return err
	}

	// Stream public posts
	for event := range publicStream {
		if updateEvent, ok := event.(*mastodon.UpdateEvent); ok {
			// Convert *mastodon.Status to *api.PublicPost
			post, err := convertToPublicPost(updateEvent.Status)
			if err != nil {
				log.Fatalf("Error converting post: %v", err)
			}

			handler(post)
		}
	}

	return nil
}

func convertToPublicPost(status *mastodon.Status) (*api.PublicPost, error) {
	var post api.PublicPost
	err := mapstructure.Decode(status, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
