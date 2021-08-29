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
	GetById(id, userId int) (hotelManager.Room, error)
	Delete(id, userId int) error
	Update(id, userId int, input hotelManager.UpdateRoomInput) error
}

type Booking interface {
	Create(roomId int, booking hotelManager.Booking) (int, error)
	GetAll(userId, roomId int) ([]hotelManager.Booking, error)
	GetById(userId, bookingId int) (hotelManager.Booking, error)
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
		Booking:       NewBookingPostgres(db),
	}
}
