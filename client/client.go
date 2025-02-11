package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

/**
予約処理
*/

/**
予約キャンセル
*/

/*
*
clientサーバを立ち上げる
*/
func main() {
	ctx := context.Background()
	projectID := "billshotelapp-450614"
	subscriptionID := "reserve-sub"

	// Pub/Sub クライアント作成
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// サブスクリプション取得
	sub := client.Subscription(subscriptionID)

	// メッセージを非同期で通信
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("📩 受信したメッセージ: %s\n", string(msg.Data))
		msg.Ack() //受信メッセージを削除
	})
	if err != nil {
		log.Fatalf("Receive failed: %v", err)
	}
}
