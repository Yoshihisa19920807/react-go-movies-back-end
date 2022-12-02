package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}