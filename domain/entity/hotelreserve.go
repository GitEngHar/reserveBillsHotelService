package entity

type HotelReserve struct {
	ID               int   `json:"id"`
	IsCancel         bool  `json:"is_cancel"`
	HotelID          int   `json:"hotel_id"`
	UserID           int   `json:"user_id"`
	ReservedDatetime int64 `json:"reserved_datetime"`
	CheckInDatetime  int64 `json:"checkin_datetime"`
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

func HotelReserveCancel() bool {
	return false
}

func DiscountRoomsAvailable(roomsAvailable int) int {
	return roomsAvailable - 1
}

func UpscountRoomsAvailable(roomsAvailable int) int {
	return roomsAvailable + 1
}
