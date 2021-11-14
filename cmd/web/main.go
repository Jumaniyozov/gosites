package main

import (
	"github.com/jumaniyozov/gosites/pkg/config"
	"github.com/jumaniyozov/gosites/pkg/handlers"
	"github.com/jumaniyozov/gosites/pkg/utils"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig

	tc, err := utils.CreateTemplateCache()
	if err != nil {
		log.Fatalf("cannot create template cache: %v", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	app.PortNumber = ":8080"

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	utils.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Application started on port %v", app.PortNumber)

	//log.Fatal(http.ListenAndServe(app.PortNumber, nil))
	srv := &http.Server{
		Addr:    app.PortNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting application %v", err)
	}
}
