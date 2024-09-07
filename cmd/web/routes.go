package main

import (
	"net/http"

	"github.com/babtun123/MelodyV2/internal/config"
	"github.com/babtun123/MelodyV2/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	_ = app

	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/play", handlers.Repo.Play)
	mux.Get("/learn", handlers.Repo.Learn)
	mux.Get("/forum", handlers.Repo.Forum)
	mux.Get("/leaderboard", handlers.Repo.Leaderboard)
	mux.Get("/profile", handlers.Repo.Profile)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
