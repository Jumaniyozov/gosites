package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jumaniyozov/gosites/pkg/config"
	"github.com/jumaniyozov/gosites/pkg/handlers"
	"github.com/jumaniyozov/gosites/pkg/middlewares"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.MiddlewareRepo.NoSurf)
	mux.Use(middlewares.MiddlewareRepo.SessionLoad)

	mux.Get("/", handlers.HandlerRepo.Home)
	mux.Get("/about", handlers.HandlerRepo.About)
	mux.Get("/generals-quarters", handlers.HandlerRepo.Generals)
	mux.Get("/majors-suite", handlers.HandlerRepo.Majors)
	mux.Get("/search-availability", handlers.HandlerRepo.Availability)
	mux.Get("/search-availability-json", handlers.HandlerRepo.AvailabilityJSON)
	mux.Post("/search-availability", handlers.HandlerRepo.PostAvailability)
	mux.Get("/make-reservation", handlers.HandlerRepo.Reservation)
	mux.Get("/reservation-summary", handlers.HandlerRepo.ReservationSummary)
	mux.Post("/make-reservation", handlers.HandlerRepo.PostReservation)
	mux.Get("/contact", handlers.HandlerRepo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
