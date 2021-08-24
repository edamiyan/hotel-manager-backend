package repository

import (
	"github.com/edamiyan/hotel-manager"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user hotelManager.User) (int, error)
	GetUser(username, password string) (hotelManager.User, error)
}

type Room interface {
	Create(userId int, room hotelManager.Room) (int, error)
	GetAll(userId int) ([]hotelManager.Room, error)
}

type Booking interface {
}

type Repository struct {
	Authorization
	Booking
	Room
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Room:          NewRoomPostgres(db),
	}
}
