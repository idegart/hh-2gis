package room

import (
	"time"
)

type repository interface {
	Lock()
	Unlock()
	CheckRoom(hotelID, roomID string, date time.Time) error
	BookRoom(hotelID, roomID string, date time.Time) error
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) BookRoom(hotelID, roomID string, from time.Time, to time.Time) error {
	s.repository.Lock()
	defer s.repository.Unlock()

	date := from

	for date.Before(to) || date.Equal(to) {
		if err := s.repository.CheckRoom(hotelID, roomID, date); err != nil {
			return err
		}
		date = date.AddDate(0, 0, 1)
	}

	date = from

	for date.Before(to) || date.Equal(to) {
		if err := s.repository.BookRoom(hotelID, roomID, date); err != nil {
			return err
		}
		date = date.AddDate(0, 0, 1)
	}

	return nil
}
