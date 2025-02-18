package usecase

import (
	"reserveBillsHotelService/domain/entity"
	"reserveBillsHotelService/domain/repository"
)

func Reserve(hotel *entity.Hotel, hotelReserve *entity.HotelReserve, repository repository.ReserveHotelRepository) *entity.Hotel {
	// ホテルが予約可能かを確認する
	if entity.CanReserve(hotel.RoomsAvailable) {
		err := repository.RegurationReserveHotel(hotelReserve)
		if err != nil {
			return nil
		} // 予約実施
		discountedRoomsAvailable := entity.DiscountRoomsAvailable(hotel.RoomsAvailable) // 減算処理
		hotel.RoomsAvailable = discountedRoomsAvailable
		return hotel
	} else {
		return nil
	}
}

func Cancel(hotel *entity.Hotel, hotelReserve *entity.HotelReserve, repository repository.ReserveHotelRepository) *entity.Hotel {
	hotelReserve.IsCancel = entity.HotelReserveCancel() // キャンセル処理
	err := repository.CancelReserveHotel(hotelReserve)  // DBキャンセル処理
	if err != nil {
		return nil
	}
	// 空き部屋数の加算処理
	hotel.RoomsAvailable = entity.UpscountRoomsAvailable(hotel.RoomsAvailable)
	return hotel
}
