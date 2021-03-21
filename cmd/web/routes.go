package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/philippmossier/golang-booking-app/pkg/config"
	"github.com/philippmossier/golang-booking-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
