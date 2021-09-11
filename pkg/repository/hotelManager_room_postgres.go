package repository

import (
	"fmt"
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type RoomPostgres struct {
	db *sqlx.DB
}

func NewRoomPostgres(db *sqlx.DB) *RoomPostgres {
	return &RoomPostgres{db: db}
}

func (r *RoomPostgres) Create(userId int, room hotelManager.Room) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createRoomQuery := fmt.Sprintf("INSERT INTO %s "+
		"(room_number, double_bed, single_bed, description, price) VALUES ($1, $2, $3, $4, $5) RETURNING id", roomsTable)
	row := tx.QueryRow(createRoomQuery,
		room.RoomNumber,
		room.DoubleBed,
		room.SingleBed,
		room.Description,
		room.Price)
	if err = row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUserRoomQuery := fmt.Sprintf("INSERT INTO %s (user_id, room_id) VALUES ($1, $2)", usersRoomsTable)
	_, err = tx.Exec(createUserRoomQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *RoomPostgres) GetAll(userId int) ([]hotelManager.Room, error) {
	var rooms []hotelManager.Room
	query := fmt.Sprintf("SELECT rt.id, rt.room_number, rt.double_bed, rt.single_bed, rt.description, rt.price"+
		" FROM %s rt INNER JOIN %s ur ON rt.id = ur.room_id WHERE ur.user_id = $1", roomsTable, usersRoomsTable)
	err := r.db.Select(&rooms, query, userId)

	return rooms, err
}

func (r *RoomPostgres) GetById(id, userId int) (hotelManager.Room, error) {
	var room hotelManager.Room
	query := fmt.Sprintf("SELECT rt.id, rt.room_number, rt.double_bed, rt.single_bed, rt.description, rt.price FROM %s rt INNER JOIN %s ur ON rt.id = ur.room_id WHERE ur.user_id = $1 AND ur.room_id = $2", roomsTable, usersRoomsTable)
	err := r.db.Get(&room, query, userId, id)

	return room, err
}

func (r *RoomPostgres) Delete(id, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s rt USING %s ur WHERE rt.id = ur.room_id AND ur.user_id = $1 AND ur.room_id = $2", roomsTable, usersRoomsTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}

func (r *RoomPostgres) Update(id, userId int, input hotelManager.UpdateRoomInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.RoomNumber != nil {
		setValues = append(setValues, fmt.Sprintf("room_number=$%d", argId))
		args = append(args, *input.RoomNumber)
		argId++
	}

	if input.SingleBed != nil {
		setValues = append(setValues, fmt.Sprintf("single_bed=$%d", argId))
		args = append(args, *input.SingleBed)
		argId++
	}

	if input.DoubleBed != nil {
		setValues = append(setValues, fmt.Sprintf("double_bed=$%d", argId))
		args = append(args, *input.DoubleBed)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s rt SET %s FROM %s ur WHERE rt.id = ur.room_id AND ur.room_id=$%d AND ur.user_id=$%d",
		roomsTable, setQuery, usersRoomsTable, argId, argId+1)
	args = append(args, id, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
