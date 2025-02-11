package main

import (
	"context"
	"log"
	"reserveBillsHotelService/infrastructure/message"
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
	ctx, cancel := context.WithCancel(context.Background())

	// Pub/Sub クライアント作成
	client, err := message.NewPubSubClient(ctx)
	if err != nil {
		log.Fatalf("Client failed: %v", err)
	}
	defer client.Close()
	// サブスクリプション取得
	recievedMessage, err := message.RevieveMessage(ctx, cancel, client)
	if err != nil {
		log.Fatalf("Receive failed: %v", err)
	}
	log.Fatalf("Success!! ✉️ : %v", recievedMessage)
}
