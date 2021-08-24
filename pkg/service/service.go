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
}

type Booking interface {
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
	}
}
