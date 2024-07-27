package order

import (
	"time"

	ds "2GIS/internal/datastruct"
)

type repository interface {
	SaveOrder(order *ds.Order) error
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateOrder(email string, hotelID, roomID string, from time.Time, to time.Time) (*ds.Order, error) {
	order := &ds.Order{
		Email:   email,
		HotelID: hotelID,
		RoomID:  roomID,
		From:    from,
		To:      to,
	}

	if err := s.repository.SaveOrder(order); err != nil {
		return nil, err
	}

	return order, nil
}
