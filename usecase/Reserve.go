package usecase

import (
	"reserveBillsHotelService/domain/entity"
	"reserveBillsHotelService/domain/repository"
)

func Reserve(hotel *entity.Hotel, hotelReserve *entity.HotelReserve, repository repository.ReserveHotelRepository) *entity.Hotel {
	// ホテルが予約可能かを確認する
	if entity.CanReserve(hotel.RoomsAvailable) {
		repository.RegurationReserveHotel(hotelReserve)                                 // 予約実施
		discountedRoomsAvailable := entity.DiscountRoomsAvailable(hotel.RoomsAvailable) // 減算処理
		hotel.RoomsAvailable = discountedRoomsAvailable
		return hotel
	} else {
		return nil
	}
}
