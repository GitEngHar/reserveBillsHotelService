package main

import (
	"reserveBillsHotelService/client/repository"
	"reserveBillsHotelService/domain/entity"
	"reserveBillsHotelService/infrastructure/database"
	"time"
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
	/* PUBSUBの実装コメントアウト
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
	log.Fatalf("Success!! ✉️ : %v", recievedMessage)*/

	// db実体化
	db, _ := database.NewMySQL()
	// 予約MYSQLの実行
	reserveHotel := repository.NewReserveHotel(db)
	// ホテル予約 実体化
	hotelReserve := entity.NewHotelReserve(0, false, 0, 1000, time.Now().Unix(), time.Now().Unix())
	//ホテル予約実行
	err := reserveHotel.RegurationReserveHotel(hotelReserve)
	if err != nil {
		panic(err)
	}
}
