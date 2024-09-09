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
	mux.Get("/match", handlers.Repo.Match)
	mux.Get("/lessonone", handlers.Repo.LessonOne)
	mux.Get("/lessontwo", handlers.Repo.LessonTwo)
	mux.Get("/piano", handlers.Repo.Piano)
	mux.Get("/quiz", handlers.Repo.Quiz)
	mux.Get("/identify", handlers.Repo.Identify)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
