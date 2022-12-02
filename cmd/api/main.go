package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string

}

func main() {
	// set application config
	// 宣言のみ。　代入なし
	var app application // application is the type that is declared above

	// read from command line

	// connect to the database

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	// start a web server
	// 宣言と代入を一緒にするパターン (var省略)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		log.Fatal(err)
	}
}