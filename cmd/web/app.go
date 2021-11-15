package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/jumaniyozov/gosites/pkg/config"
	"github.com/jumaniyozov/gosites/pkg/handlers"
	"github.com/jumaniyozov/gosites/pkg/helpers"
	"github.com/jumaniyozov/gosites/pkg/middlewares"
	"github.com/jumaniyozov/gosites/pkg/models"
	"github.com/jumaniyozov/gosites/pkg/utils"
	"log"
	"net/http"
	"os"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// InitializeApp initializes application instance
func InitializeApp() *config.AppConfig {
	gob.Register(models.Reservation{})

	// is app in production or in development
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := utils.CreateTemplateCache()
	if err != nil {
		log.Fatalf("cannot create template cache: %v", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.PortNumber = ":8080"

	handlersRepo := handlers.NewRepo(&app)
	handlers.NewHandlers(handlersRepo)
	utils.NewTemplates(&app)
	helpers.NewHelpers(&app)

	middlewareRepo := middlewares.NewMiddlewareRepo(&app)
	middlewares.NewMiddlewares(middlewareRepo)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Application started on port %v", app.PortNumber)

	return &app
}
