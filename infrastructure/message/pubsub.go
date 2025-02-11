package message

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

const (
	projectID      = "billshotelapp-450614"
	subscriptionID = "reserve-sub"
)

func NewPubSubClient(ctx context.Context) (*pubsub.Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}

func RevieveMessage(ctx context.Context, cancel context.CancelFunc, client *pubsub.Client) (string, error) {
	var returnMessage = ""
	sub := client.Subscription(subscriptionID)
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("📩 受信したメッセージ: %s\n", string(msg.Data))
		returnMessage = string(msg.Data)
		msg.Ack()
		cancel()
	})
	if err != nil {
		log.Fatalf("Failed to recieve message : %v", err)
		return "", err
	}
	return returnMessage, nil
}
