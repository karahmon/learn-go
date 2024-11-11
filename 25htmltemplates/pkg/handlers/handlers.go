package handlers

import (
	"main/25htmltemplates/pkg/config"
	"main/25htmltemplates/pkg/models"
	"main/25htmltemplates/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.Appconfig
}

// Creates a ne new repository
func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets Repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,Again."
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap})
}
