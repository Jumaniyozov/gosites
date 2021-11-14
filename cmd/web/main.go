package main

import (
	"github.com/jumaniyozov/gosites/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Application started on port %v", portNumber)

	log.Fatal(http.ListenAndServe(portNumber, nil))

}
