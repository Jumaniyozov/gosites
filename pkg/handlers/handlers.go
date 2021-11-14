package handlers

import (
	"github.com/jumaniyozov/gosites/pkg/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "about.page.tmpl")
}
