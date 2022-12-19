package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data.omitempty"`
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	var out []byte
	var err error

	// json.Marshal() returns nil if the argument is an empty slice
	if reflect.ValueOf(data).Len() == 0 {
		out, err = json.Marshal([]string{})
	} else {
		out, err = json.Marshal(data)
	}

	// out, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}
