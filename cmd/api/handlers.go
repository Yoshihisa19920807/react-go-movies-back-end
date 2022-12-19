package main

import (
	"backend/internal/models" // root directory is defined as "backend" which is declared in go.mod
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world from %s", app.Domain)
	// var <変数> = <型定義><実際のキーバリュー>
	var payload = struct {
		Status string `json:"status"`; // <キー> <型> <デフォルト>
		Message string `json:"message"`;
		Version string `json:"version"`
	}{
		Status: "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}
	// ↑は
	// hoge_struct := struct {
	//   Status string `json:"status"`; // <キー> <型> <デフォルト>
	//   Message string `json:"message"`;
	//   Version string `json:"version"`
	// }
	// var payload = hoge_struct{
	// 	Status: "active",
	// 	Message: "Go Movies up and running",
	// 	Version: "1.0.0",
	// }
	// と同じ

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	// time.Parse(<format or example>, <value>)
	rd, _ := time.Parse("2006-01-02", "2010-01-09")
	fmt.Print(rd)
	_500DaysOfSummer := models.Movie {
		Id: 1,
		Title: "500 days of Summer",
		ReleaseDate: rd,
		MPAARating: "PG-13",
		RunTime: 95,
		Description: "Romance which even men can enjoy.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	movies = append(movies, _500DaysOfSummer)

	rd, _ = time.Parse("2006-01-02", "2009-03-20")
	YesMan:= models.Movie {
		Id: 2,
		Title: "Yes Man",
		ReleaseDate: rd,
		MPAARating: "PG-13",
		RunTime: 104,
		Description: "Comedy that will encourage you to take a new step.",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	movies = append(movies, YesMan)
	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
