package message

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
)

const (
	projectID = "billshotelapp-450614"
)

func NewPubSubClient(ctx context.Context) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}
