package order_creation

import (
	"time"

	ds "2GIS/internal/datastruct"
)

type Input struct {
	HotelID string
	RoomID  string
	Email   string
	From    time.Time
	To      time.Time
}

type Output struct {
	Order *ds.Order
}
