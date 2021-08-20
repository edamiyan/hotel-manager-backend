package service

import "github.com/edamiyan/hotel-manager/pkg/repository"

type Authorization interface {
}

type Room interface {
}

type Booking interface {
}

type Service struct {
	Authorization
	Booking
	Room
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
