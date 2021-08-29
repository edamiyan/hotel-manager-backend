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

func (s *RoomService) GetById(id, userId int) (hotelManager.Room, error) {
	return s.repo.GetById(id, userId)
}

func (s *RoomService) Delete(id, userId int) error {
	return s.repo.Delete(id, userId)
}

func (s *RoomService) Update(id int, userId int, input hotelManager.UpdateRoomInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, userId, input)
}
