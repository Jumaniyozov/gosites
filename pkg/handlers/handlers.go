package handlers

import (
	"github.com/jumaniyozov/gosites/pkg/config"
	"github.com/jumaniyozov/gosites/pkg/models"
	"github.com/jumaniyozov/gosites/pkg/utils"
	"net/http"
)

var HandlerRepo *HandlerRepository

type HandlerRepository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *HandlerRepository {
	return &HandlerRepository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *HandlerRepository) {
	HandlerRepo = r
}

func (m *HandlerRepository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	utils.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *HandlerRepository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIp

	utils.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
