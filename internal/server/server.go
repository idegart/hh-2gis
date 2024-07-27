package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type controller interface {
	PostOrders(w http.ResponseWriter, r *http.Request)
}

type App struct {
	router chi.Router
}

func NewApp(c controller) *App {
	router := chi.NewRouter()

	router.Post("/orders", c.PostOrders)

	return &App{
		router: router,
	}
}

func (a *App) Run(port string) error {
	return http.ListenAndServe(port, a.router)
}
