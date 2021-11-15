package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jumaniyozov/gosites/pkg/config"
	"github.com/jumaniyozov/gosites/pkg/forms"
	"github.com/jumaniyozov/gosites/pkg/helpers"
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
	utils.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *HandlerRepository) About(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *HandlerRepository) Generals(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *HandlerRepository) Majors(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *HandlerRepository) Availability(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"messages"`
}

func (m *HandlerRepository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.Marshal(resp)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}

func (m *HandlerRepository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	_, err := w.Write([]byte(fmt.Sprintf("Start date is %v and end date is %v", start, end)))
	if err != nil {
		helpers.ServerError(w, err)
	}
}

func (m *HandlerRepository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	utils.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		//Data: data,
	})
}

func (m *HandlerRepository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)
	//form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		utils.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *HandlerRepository) Contact(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *HandlerRepository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	utils.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
