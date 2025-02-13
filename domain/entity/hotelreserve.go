package entity

type HotelReserve struct {
	ID               int
	IsCancel         bool
	HotelID          int
	UserID           int
	ReservedDatetime int64
	CheckInDatetime  int64
}

func NewHotelReserve(id int, isCancel bool, hotelID int, userID int, reservedDatetime int64, CheckInDatetime int64) *HotelReserve {
	return &HotelReserve{
		ID:               id,
		IsCancel:         isCancel,
		HotelID:          hotelID,
		UserID:           userID,
		ReservedDatetime: reservedDatetime,
		CheckInDatetime:  CheckInDatetime,
	}
}

func CanReserve(roomsAvailable int) bool {
	if roomsAvailable > 0 {
		return true
	} else {
		return false
	}
}

func DiscountRoomsAvailable(roomsAvailable int) int {
	return roomsAvailable - 1
}
