package order_creation

import (
	"time"

	ds "2GIS/internal/datastruct"
)

type roomService interface {
	BookRoom(hotelID, roomID string, from time.Time, to time.Time) error
}

type orderService interface {
	CreateOrder(email string, hotelID, roomID string, from time.Time, to time.Time) (*ds.Order, error)
}

type UseCase struct {
	roomService  roomService
	orderService orderService
}

func NewUseCase(roomService roomService, orderService orderService) *UseCase {
	return &UseCase{
		roomService:  roomService,
		orderService: orderService,
	}
}

func (u *UseCase) Handle(input Input) (*Output, error) {
	if err := u.roomService.BookRoom(input.HotelID, input.RoomID, input.From, input.To); err != nil {
		return nil, err
	}

	order, err := u.orderService.CreateOrder(input.Email, input.HotelID, input.RoomID, input.From, input.To)
	if err != nil {
		return nil, err
	}

	return &Output{Order: order}, nil
}
