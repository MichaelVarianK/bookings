package handlers

import (
	"net/http"

	"github.com/MichaelVarianK/bookings/pkg/config"
	"github.com/MichaelVarianK/bookings/pkg/models"
	"github.com/MichaelVarianK/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform Some Business Logic
	msg := make(map[string]string)
	msg["test"] = "Hello, World!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	msg["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: msg,
	})
}