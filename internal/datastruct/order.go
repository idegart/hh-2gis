package datastruct

import "time"

type Order struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}
