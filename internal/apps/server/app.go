package server

import (
	"log/slog"

	"2GIS/internal/controller"
	"2GIS/internal/repository/inmemory"
	"2GIS/internal/server"
	"2GIS/internal/services/order"
	"2GIS/internal/services/room"
	order_creation "2GIS/internal/usecases/order-creation"
)

func Run() {
	// DB
	repo := inmemory.NewRepository()
	repo.FillDemo()

	// Services
	roomService := room.NewService(repo)
	orderService := order.NewService(repo)

	// Use-Cases
	orderCreation := order_creation.NewUseCase(roomService, orderService)

	// Controller
	ctr := controller.NewController(orderCreation)

	// Server
	srv := server.NewApp(ctr)

	slog.Info("starting app")
	slog.Error("app error", srv.Run(":8000"))
}
