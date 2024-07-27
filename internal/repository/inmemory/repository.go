package inmemory

import (
	"fmt"
	"sync"
	"time"

	ds "2GIS/internal/datastruct"
	"github.com/google/uuid"
)

type Repository struct {
	mx sync.Mutex

	orders       map[string]ds.Order
	availability map[string]map[string]map[time.Time]int // hotel -> room -> time -> quota
}

func NewRepository() *Repository {
	return &Repository{
		orders:       make(map[string]ds.Order),
		availability: make(map[string]map[string]map[time.Time]int),
	}
}

func (r *Repository) Lock() {
	r.mx.Lock()
}

func (r *Repository) Unlock() {
	r.mx.Unlock()
}

func (r *Repository) RoomExists(hotelID, roomID string) bool {
	if _, ok := r.availability[hotelID]; !ok {
		return false
	}

	if _, ok := r.availability[hotelID][roomID]; !ok {
		return false
	}

	return true
}

func (r *Repository) BookRoom(hotelID, roomID string, date time.Time) error {
	r.availability[hotelID][roomID][date]--
	return nil
}

func (r *Repository) CheckRoom(hotelID, roomID string, date time.Time) error {
	if q, ok := r.availability[hotelID][roomID][date]; !ok {
		return fmt.Errorf("room not available for date: %v", date)
	} else if q < 1 {
		return fmt.Errorf("room not have enouhg quota for date: %v", date)
	}

	return nil
}

func (r *Repository) SaveOrder(order *ds.Order) error {
	id := uuid.New().String()

	order.ID = id
	r.orders[id] = *order

	return nil
}

func (r *Repository) FillDemo() {
	const hotel = "demo"
	const room1 = "demo-1"
	const room2 = "demo-2"

	r.availability = map[string]map[string]map[time.Time]int{
		hotel: {
			room1: {
				toDate(2024, 1, 1): 1,
				toDate(2024, 1, 2): 2,
				toDate(2024, 1, 3): 1,
			},
			room2: {
				toDate(2024, 1, 1): 2,
				toDate(2024, 1, 2): 1,
				toDate(2024, 1, 3): 3,
			},
		},
	}
}

func toDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
