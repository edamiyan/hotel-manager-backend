package service

import (
	"github.com/edamiyan/hotel-manager"
	"github.com/edamiyan/hotel-manager/pkg/repository"
)

type Authorization interface {
	CreateUser(user hotelManager.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Room interface {
	Create(userId int, room hotelManager.Room) (int, error)
	GetAll(userId int) ([]hotelManager.Room, error)
	GetById(id, userId int) (hotelManager.Room, error)
	Delete(id, userId int) error
	Update(id, userId int, input hotelManager.UpdateRoomInput) error
}

type Booking interface {
	Create(userId int, roomId int, booking hotelManager.Booking) (int, error)
	GetAll(userId, roomId int) ([]hotelManager.Booking, error)
	GetById(userId, bookingId int) (hotelManager.Booking, error)
	Update(userId, bookingId int, input hotelManager.UpdateBookingInput) error
	Delete(userId, bookingId int) error
	GetRoomIdByBooking(userId, bookingid int) (int, error)
	GetAllUserBookings(userId int) ([]hotelManager.BookingTimeline, error)
}

type Service struct {
	Authorization
	Booking
	Room
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Room:          NewRoomService(repos.Room),
		Booking:       NewBookingService(repos.Booking, repos.Room),
	}
}
