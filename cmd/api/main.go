package main

import (
	"database/sql"
	"flag"
	"log"
)

const port = 8080

type application struct {
	Domain string
	DSN string // data string name
	DB *sql.DB
}

func main() {
	// set application config
	// 宣言のみ。　代入なし
	var app application // application is the type that is declared above

	// // read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "postgres connection string")
	flag.Parse()

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = conn

	app.Domain = "example.com"

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	// // start a web server
	// // 宣言と代入を一緒にするパターン (var省略)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	// if err != nil {
	// 	log.Fatal(err)
	// }
}