package handlers

import (
	"github.com/jumaniyozov/gosites/pkg/config"
	"github.com/jumaniyozov/gosites/pkg/models"
	"github.com/jumaniyozov/gosites/pkg/utils"
	"net/http"
)

var Repo *Repository

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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, there."

	utils.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
