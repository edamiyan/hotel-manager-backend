package service

import (
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/edamiyan/hotel-manager/pkg/repository"
)

type RoomService struct {
	repo repository.Room
}

func NewRoomService(repo repository.Room) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) Create(userId int, room hotelManager.Room) (int, error) {
	return s.repo.Create(userId, room)
}

func (s *RoomService) GetAll(userId int) ([]hotelManager.Room, error) {
	return s.repo.GetAll(userId)
}
