package repository

import (
	"database/sql"
	"reserveBillsHotelService/client/repository"
)

type ReserveHotelRepository interface {
	NewHotelReserveRepository(db *sql.DB) *repository.ReserveHotel
	RegurationReserveHotel()
	CancelReserveHotel()
}
