package repository

import (
	"database/sql"
	"reserveBillsHotelService/domain/entity"
)

type ReserveHotel struct {
	db *sql.DB
}

func NewReserveHotel(db *sql.DB) *ReserveHotel {
	return &ReserveHotel{db: db}
}

func (r *ReserveHotel) RegurationReserveHotel(hotel *entity.HotelReserve) error {
	_, err := r.db.Exec(
		"INSERT INTO reserve_hotel (id,iscancel,hotelid,userid,reserved_unix_datetime,checkin_unix_datetime) VALUES(?,?,?,?,?,?)",
		hotel.ID, hotel.IsCancel, hotel.HotelID, hotel.UserID, hotel.ReservedDatetime, hotel.CheckInDatetime,
	)
	return err
}
