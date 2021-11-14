package main

import (
	"log"
	"net/http"
)

func main() {
	app := InitializeApp()

	//log.Fatal(http.ListenAndServe(app.PortNumber, nil))
	srv := &http.Server{
		Addr:    app.PortNumber,
		Handler: routes(app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while starting application %v", err)
	}
}
