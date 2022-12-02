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

	// when someone calls listenAndServe it calls port 8080 on our web server and they're looking for just the root level of the application and execute "Hello" function
	http.HandleFunc("/", Hello)

	// start a web server
	// 宣言と代入を一緒にするパターン (var省略)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatal(err)
	}
}