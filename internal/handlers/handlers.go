package handlers

import (
	"net/http"

	"github.com/babtun123/MelodyV2/internal/config"
	"github.com/babtun123/MelodyV2/internal/models"
	"github.com/babtun123/MelodyV2/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo Creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// Play is the play page handler
func (m *Repository) Play(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "play.page.tmpl", &models.TemplateData{})
}

// Learn is the learn page handler
func (m *Repository) Learn(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "learn.page.tmpl", &models.TemplateData{})
}

// Forum is the forum page handler
func (m *Repository) Forum(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "forum.page.tmpl", &models.TemplateData{})
}

// Leaderboard is the leaderboard page handler
func (m *Repository) Leaderboard(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "leaderboard.page.tmpl", &models.TemplateData{})
}

// Profile is the profile page handler
func (m *Repository) Profile(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "profile.page.tmpl", &models.TemplateData{})
}

// Match is the match page handler
func (m *Repository) Match(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "matchmenu.page.tmpl", &models.TemplateData{})
}

// LessonOne is the lesson 1 page handler
func (m *Repository) LessonOne(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "lesson1.page.tmpl", &models.TemplateData{})
}

// LessonTwo is the LessonTwo page handler
func (m *Repository) LessonTwo(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "lesson2.page.tmpl", &models.TemplateData{})
}

// Piano is the piano page handler
func (m *Repository) Piano(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "piano.page.tmpl", &models.TemplateData{})
}

// Quiz is the quiz page handler
func (m *Repository) Quiz(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "quiz.page.tmpl", &models.TemplateData{})
}

// Identify is the identify page handler
func (m *Repository) Identify(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "identify.page.tmpl", &models.TemplateData{})
}
