package repository

import (
	"database/sql"
)

type ReserveHotel struct {
	db *sql.DB
}

func NewReserveHotel(db *sql.DB) *ReserveHotel {
	return &ReserveHotel{db: db}
}

func (r *ReserveHotel) RegurationReserveHotel(hotelID int, userID int) {
	//entity
}
func CancelReserveHotel() {}
