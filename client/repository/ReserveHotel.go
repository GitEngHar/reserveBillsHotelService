package repository

import (
	"database/sql"
	"reserveBillsHotelService/domain/entity"
	"reserveBillsHotelService/domain/repository"
)

type ReserveHotel struct {
	db *sql.DB
}

func NewHotelReserveRepository(db *sql.DB) repository.ReserveHotelRepository {
	return &ReserveHotel{db: db}
}

func (r *ReserveHotel) RegurationReserveHotel(hotel *entity.HotelReserve) error {
	_, err := r.db.Exec(
		"INSERT INTO reserve_hotel (id,iscancel,hotelid,userid,reserved_unix_datetime,checkin_unix_datetime) VALUES(?,?,?,?,?,?)",
		hotel.ID, hotel.IsCancel, hotel.HotelID, hotel.UserID, hotel.ReservedDatetime, hotel.CheckInDatetime,
	)
	return err
}

func (r *ReserveHotel) CancelReserveHotel(hotel *entity.HotelReserve) error {
	_, err := r.db.Exec(
		"UPDATE reserve_hotel SET iscancel=(?) WHERE id=(?)",
		hotel.IsCancel, hotel.ID,
	)
	return err
}
