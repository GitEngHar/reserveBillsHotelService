package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
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
	"strconv"
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
func (s *Subscriber) Receive(ctx context.Context, handler func(context.Context, *pubsub.Message) error) {
	err := s.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("📩 受信したメッセージ: %s\n", string(msg.Data))
		defer msg.Ack() //関数が終了した時にmsgをackする
		err := handler(ctx, msg)
		if err != nil {
			log.Fatalf("failed message process %v", err)
		}
	})
	if err != nil {
		log.Fatalf("failed message revieve%v", err)
	}
}

func (s *Subscriber) grpcHandler(ctx context.Context, msg *pubsub.Message) error {
	//ホテル一覧の取得
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	//msgを解析してNewHotelへ代入する
	var hotelReserve *entity.HotelReserve
	err := json.Unmarshal(msg.Data, &hotelReserve)
	if err != nil {
		msg.Nack()
		log.Fatalf("JSON の解析に失敗しました:%v", err)
		return err
	}

	//一旦MockでNewHotelを生成する
	//hotelReserve := entity.NewHotelReserve(0, false, 0, 1000, time.Now().Unix(), time.Now().Unix())
	//Hotel情報を取得する
	resp, err := s.client.GetHotel(ctx, &pb.HotelRequest{Id: int32(hotelReserve.HotelID)})
	if err != nil {
		msg.Nack()
		log.Fatalf("Server Error : %v", err)
	}
	rsvHotel := resp.GetHotel()
	hotel := entity.NewHotel(int(rsvHotel.Id), rsvHotel.Name, int(rsvHotel.PricePernight), int(rsvHotel.RoomsAvailable))
	fmt.Println("GetHotel 🏨 ID: ", hotel.ID)
	hotel = usecase.Reserve(hotel, hotelReserve, s.repository)
	if hotel != nil {
		//ホテルの予約ができた場合空室を-1にして返す
		resp, err = s.client.UpdateHotel(ctx, &pb.HotelRequest{
			Id:             int32(hotel.ID),
			Name:           hotel.Name,
			PricePernight:  int32(hotel.PricePerNight),
			RoomsAvailable: int32(hotel.RoomsAvailable),
		})
		if err != nil {
			msg.Nack()
			log.Fatalf("Server Error : %v", err)
			return err
		}
		fmt.Printf("UserID:%v のお客様の予約が完了しました🎉\n", strconv.Itoa(hotelReserve.UserID))
	} else {
		//ホテルの予約ができない場合nilにして返す
		log.Fatalln("満室で予約ができませんでした😭")
	}
	return nil
}

/*
*
clientサーバを立ち上げる
*/
func main() {
	ctx := context.Background()
	// hotel-server-service
	connection, err := grpc.NewClient("hotel-server-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
	sub := pubSubClient.Subscription("reserve-sub-dev")
	subscriber := &Subscriber{sub: sub, client: pb.NewHotelServiceClient(connection), repository: reserveHotel}
	subscriber.Receive(ctx, subscriber.grpcHandler)

}
