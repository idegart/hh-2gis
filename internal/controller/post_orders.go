package controller

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	order_creation "2GIS/internal/usecases/order-creation"
)

const opPostOrders = "[POST ORDERS] "

type postOrdersRequest struct {
	HotelID string    `json:"hotel_id"`
	RoomID  string    `json:"room_id"`
	Email   string    `json:"email"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}

func (c *Controller) PostOrders(w http.ResponseWriter, r *http.Request) {
	slog.Info(opPostOrders + "handle post orders request")

	decoder := json.NewDecoder(r.Body)

	var data postOrdersRequest

	if err := decoder.Decode(&data); err != nil {
		if errors.Is(err, io.EOF) {
			slog.Error(opPostOrders + "handle empty request")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("empty request"))
			return
		}
		slog.Error(opPostOrders+"decode data", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	if err := validatePostOrdersData(data); err != nil {
		slog.Error(opPostOrders+"validation error", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	out, err := c.orderCreation.Handle(order_creation.Input{
		HotelID: data.HotelID,
		RoomID:  data.RoomID,
		Email:   data.Email,
		From:    data.From,
		To:      data.To,
	})
	if err != nil {
		slog.Error("order creation error", "err", err)
		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(out.Order)
	_, _ = w.Write(response)
}

func validatePostOrdersData(data postOrdersRequest) error {
	if data.HotelID == "" {
		return errors.New("hotel id is empty")
	}

	if data.RoomID == "" {
		return errors.New("room id is empty")
	}

	if data.Email == "" {
		return errors.New("email is empty")
	}

	if data.From.IsZero() {
		return errors.New("from is empty")
	}

	if data.To.IsZero() {
		return errors.New("to is empty")
	}

	if data.To.Before(data.From) {
		return errors.New("to is before from")
	}

	return nil
}
