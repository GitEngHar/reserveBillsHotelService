package entity

import (
	"time"
)

type HotelReserve struct {
	ID               int
	IsCancel         bool
	HotelID          int
	UserID           int
	ReservedDatetime time.Time
	CheckInDatetime  time.Time
}
