package repository

import (
	"reserveBillsHotelService/domain/entity"
)

type ReserveHotelRepository interface {
	RegurationReserveHotel(hotelReserve *entity.HotelReserve) error
	CancelReserveHotel(hotelReserve *entity.HotelReserve) error
}
