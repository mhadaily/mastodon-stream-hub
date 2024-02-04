package util

import (
	api "github.com/mhadaily/mastodon-stream-hub/pkg/api"

	"google.golang.org/protobuf/proto"
)

// SerializePublicPost serializes a PublicPost to bytes
func SerializePublicPost(post *api.PublicPost) ([]byte, error) {
	return proto.Marshal(post)
}

// DeserializePublicPost deserializes bytes into a PublicPost
func DeserializePublicPost(data []byte) (*api.PublicPost, error) {
	post := &api.PublicPost{}
	if err := proto.Unmarshal(data, post); err != nil {
		return nil, err
	}
	return post, nil
}

// CreateGenericMessage creates a new GenericMessage for a given PublicPost
//
// The key is the post ID and the value is the serialized PublicPost
func CreateGenericMessage(post *api.PublicPost) (*api.PubSubGenericMessage, error) {
	data, err := SerializePublicPost(post)
	if err != nil {
		return nil, err
	}
	return &api.PubSubGenericMessage{
		Key:   []byte(post.Id),
		Value: data,
	}, nil
}
