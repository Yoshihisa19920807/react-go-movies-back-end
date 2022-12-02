package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8000

type application struct {
	Domain string

}

func main() {
	// set application config
	var app application // application is the type that is declared above

	// read from command line

	// connect to the database

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatal(err)
	}
}