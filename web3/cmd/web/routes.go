package main

import (
	"net/http"
	"web3/pkg/config"
	"web3/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	return mux
}