package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"reserveBillsHotelService/client/repository"
	"reserveBillsHotelService/domain/entity"
	domainrep "reserveBillsHotelService/domain/repository"
	"reserveBillsHotelService/infrastructure/database"
	"reserveBillsHotelService/infrastructure/message"
	pb "reserveBillsHotelService/proto"
	"reserveBillsHotelService/usecase"
	"time"
)

type Subscriber struct {
	sub        *pubsub.Subscription
	client     pb.HotelServiceClient
	repository domainrep.ReserveHotelRepository
}

/*
*
handlerの実装(受け取ったメッセージに応じて処理を行う)
*/
func (s *Subscriber) Receive(ctx context.Context, handler func(context.Context, string) error) {
	err := s.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("📩 受信したメッセージ: %s\n", string(msg.Data))
		defer msg.Ack() //関数が終了した時にmsgをackする
		err := handler(ctx, string(msg.Data))
		if err != nil {
			log.Fatalf("failed message process %v", err)
		}
	})
	if err != nil {
		log.Fatalf("failed message revieve%v", err)
	}
}

func (s *Subscriber) grpcHandler(ctx context.Context, msg string) error {
	//ホテル一覧の取得
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//msgを解析してNewHotelへ代入する
	//一旦MockでNewHotelを生成する
	hotelReserve := entity.NewHotelReserve(0, false, 0, 1000, time.Now().Unix(), time.Now().Unix())
	//Hotel情報を取得する
	resp, err := s.client.GetHotel(ctx, &pb.HotelRequest{Id: 0})
	if err != nil {
		panic(err)
	}
	rsvHotel := resp.GetHotel()
	hotel := entity.NewHotel(int(rsvHotel.Id), rsvHotel.Name, int(rsvHotel.PricePernight), int(rsvHotel.RoomsAvailable))
	fmt.Println("GetHotel 🏨", hotel.ID)
	usecase.Reserve(hotel, hotelReserve, s.repository)
	return nil
}

/*
*
clientサーバを立ち上げる
*/
func main() {
	ctx := context.Background()
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not : %v", err)
	}
	defer connection.Close()

	// db実体化
	db, _ := database.NewMySQL()
	// 予約MYSQLの実行
	reserveHotel := repository.NewHotelReserveRepository(db)

	// Pub/Sub クライアント作成
	pubSubClient, err := message.NewPubSubClient(ctx)
	if err != nil {
		log.Fatalf("Client failed: %v", err)
	}

	defer pubSubClient.Close()
	// サブスクリプション取得とハンドリング
	sub := pubSubClient.Subscription("reserve-sub")
	subscriber := &Subscriber{sub: sub, client: pb.NewHotelServiceClient(connection), repository: reserveHotel}
	subscriber.Receive(ctx, subscriber.grpcHandler)

}
