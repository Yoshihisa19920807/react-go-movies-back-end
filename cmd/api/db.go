package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn" // _ avoids error even if we don't use the package
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// indicates it returns 2 values
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err // it returns 2 values as the return value structure says
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
// (app *application) part is the reciever definition
func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DSN)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database") // print lne
	return connection, nil
}