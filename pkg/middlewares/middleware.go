package middlewares

import (
	"fmt"
	"github.com/jumaniyozov/gosites/pkg/config"
	"github.com/justinas/nosurf"
	"net/http"
)

var MiddlewareRepo *MiddlewareRepository

type MiddlewareRepository struct {
	App *config.AppConfig
}

// NewMiddlewareRepo creates a new middleware repository
func NewMiddlewareRepo(a *config.AppConfig) *MiddlewareRepository {
	return &MiddlewareRepository{
		App: a,
	}
}

// NewMiddlewares sets the repository for the handlers
func NewMiddlewares(m *MiddlewareRepository) {
	MiddlewareRepo = m
}

func (m *MiddlewareRepository) WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds csrf token to all POST requests
func (m *MiddlewareRepository) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   m.App.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func (m *MiddlewareRepository) SessionLoad(next http.Handler) http.Handler {
	return m.App.Session.LoadAndSave(next)
}
