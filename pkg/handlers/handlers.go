package handlers

import (
	"net/http"

	"github.com/PrivanzaWisnu/bookings/pkg/config"
	"github.com/PrivanzaWisnu/bookings/pkg/models"
	"github.com/PrivanzaWisnu/bookings/pkg/render"
)

// Used by Handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About Page Handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, it's..."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")


	stringMap["remote_ip"] = remoteIP 

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
