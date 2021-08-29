package repository

import (
	"fmt"
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/jmoiron/sqlx"
)

type BookingPostgres struct {
	db *sqlx.DB
}

func NewBookingPostgres(db *sqlx.DB) *BookingPostgres {
	return &BookingPostgres{db: db}
}

func (r *BookingPostgres) Create(roomId int, booking hotelManager.Booking) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var bookingId int
	createBookingQuery := fmt.Sprintf("INSERT INTO %s (name, phone, arrival_date, departure_date, guests_number, is_booking, comment, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", bookingsTable)
	row := tx.QueryRow(createBookingQuery,
		booking.Name,
		booking.Phone,
		booking.ArrivalDate,
		booking.DepartureDate,
		booking.GuestsNumber,
		booking.IsBooking,
		booking.Comment,
		booking.Status)
	err = row.Scan(&bookingId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createRoomBookingsQuery := fmt.Sprintf("INSERT INTO %s (room_id, booking_id) values ($1, $2)", roomsBookingsTable)
	_, err = tx.Exec(createRoomBookingsQuery, roomId, bookingId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return bookingId, tx.Commit()
}

func (r *BookingPostgres) GetAll(userId, roomId int) ([]hotelManager.Booking, error) {
	var bookings []hotelManager.Booking
	getAllBookingsQuery := fmt.Sprintf("SELECT tb.id, tb.name, tb.phone, tb.arrival_date, tb.departure_date, tb.guests_number, tb.is_booking, tb.comment, tb.status "+
		"FROM %s tb INNER JOIN %s rb on rb.booking_id = tb.id INNER JOIN %s ur on ur.room_id = rb.room_id WHERE rb.room_id = $1 AND ur.user_id = $2", bookingsTable, roomsBookingsTable, usersRoomsTable)
	if err := r.db.Select(&bookings, getAllBookingsQuery, roomId, userId); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingPostgres) GetById(userId, bookingId int) (hotelManager.Booking, error) {
	var bookings hotelManager.Booking
	getBookingQuery := fmt.Sprintf("SELECT tb.id, tb.name, tb.phone, tb.arrival_date, tb.departure_date, tb.guests_number, tb.is_booking, tb.comment, tb.status "+
		"FROM %s tb INNER JOIN %s rb on rb.booking_id = tb.id INNER JOIN %s ur on ur.room_id = rb.room_id WHERE rb.booking_id = $1 AND ur.user_id = $2", bookingsTable, roomsBookingsTable, usersRoomsTable)
	if err := r.db.Get(&bookings, getBookingQuery, bookingId, userId); err != nil {
		return bookings, err
	}

	return bookings, nil
}
