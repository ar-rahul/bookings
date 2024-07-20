package handlers

import (
	"net/http"

	"github.com/ar-rahul/bookings/pkg/config"
	"github.com/ar-rahul/bookings/pkg/models"
	"github.com/ar-rahul/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About renders the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Kings renders the kings suite booking page
func (m *Repository) Kings(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "kingsuite.page.tmpl", &models.TemplateData{})
}

// Queens renders the queen suite booking page
func (m *Repository) Queens(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "queensuite.page.tmpl", &models.TemplateData{})
}

// Reservation renders the reservation page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation renders the contact page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{})
}
