package service

import (
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/edamiyan/hotel-manager/pkg/repository"
)

type BookingService struct {
	repo     repository.Booking
	roomRepo repository.Room
}

func NewBookingService(repo repository.Booking, roomRepo repository.Room) *BookingService {
	return &BookingService{repo: repo, roomRepo: roomRepo}
}

func (s *BookingService) Create(userId, roomId int, booking hotelManager.Booking) (int, error) {
	_, err := s.roomRepo.GetById(roomId, userId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(roomId, booking)
}

func (s *BookingService) GetAll(userId, roomId int) ([]hotelManager.Booking, error) {
	_, err := s.roomRepo.GetById(roomId, userId)
	if err != nil {
		return nil, err
	}
	return s.repo.GetAll(userId, roomId)
}

func (s *BookingService) GetById(userId, bookingId int) (hotelManager.Booking, error) {
	return s.repo.GetById(userId, bookingId)
}

func (s *BookingService) Update(userId, bookingId int, input hotelManager.UpdateBookingInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, bookingId, input)
}

func (s *BookingService) Delete(userId, bookingId int) error {
	return s.repo.Delete(userId, bookingId)
}

func (s *BookingService) GetRoomIdByBooking(userId, bookingid int) (int, error) {
	return s.repo.GetRoomIdByBooking(userId, bookingid)
}

func (s *BookingService) GetAllUserBookings(userId int) ([]hotelManager.BookingTimeline, error) {
	return s.repo.GetAllUserBookings(userId)
}