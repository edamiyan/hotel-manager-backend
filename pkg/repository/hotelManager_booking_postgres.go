package repository

import (
	"fmt"
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (r *BookingPostgres) Update(userId, bookingId int, input hotelManager.UpdateBookingInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}

	if input.ArrivalDate != nil {
		setValues = append(setValues, fmt.Sprintf("arrival_date=$%d", argId))
		args = append(args, *input.ArrivalDate)
		argId++
	}

	if input.DepartureDate != nil {
		setValues = append(setValues, fmt.Sprintf("departure_date=$%d", argId))
		args = append(args, *input.DepartureDate)
		argId++
	}

	if input.GuestsNumber != nil {
		setValues = append(setValues, fmt.Sprintf("guests_number=$%d", argId))
		args = append(args, *input.GuestsNumber)
		argId++
	}

	if input.IsBooking != nil {
		setValues = append(setValues, fmt.Sprintf("is_booking=$%d", argId))
		args = append(args, *input.IsBooking)
		argId++
	}

	if input.Comment != nil {
		setValues = append(setValues, fmt.Sprintf("comment=$%d", argId))
		args = append(args, *input.Comment)
		argId++
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args = append(args, *input.Status)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s bt SET %s FROM %s rb, %s ur WHERE bt.id = rb.booking_id AND rb.room_id = ur.room_id AND ur.user_id = $%d AND bt.id = $%d",
		bookingsTable, setQuery, roomsBookingsTable, usersRoomsTable, argId, argId+1)
	args = append(args, userId, bookingId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BookingPostgres) Delete(userId, bookingId int) error {
	query := fmt.Sprintf(`DELETE FROM %s bt USING %s rb, %s ur WHERE bt.id = rb.booking_id AND rb.room_id = ur.room_id AND ur.user_id = $1 AND bt.id = $2`, bookingsTable, roomsBookingsTable, usersRoomsTable)
	_, err := r.db.Exec(query, userId, bookingId)
	return err
}

func (r *BookingPostgres) GetRoomIdByBooking(userId, bookingId int) (int, error) {
	roomId := 0
	query := fmt.Sprintf(`SELECT rb.room_id FROM %s rb INNER JOIN %s ur on rb.room_id = ur.room_id WHERE rb.booking_id = $1 AND ur.user_id = $2`, roomsBookingsTable, usersRoomsTable)
	if err := r.db.Get(&roomId, query, bookingId, userId); err != nil {
		return roomId, err
	}

	return roomId, nil
}
