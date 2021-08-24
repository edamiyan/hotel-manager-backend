package repository

import (
	"fmt"
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/jmoiron/sqlx"
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
