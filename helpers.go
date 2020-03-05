package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *server) respond(w http.ResponseWriter, r *http.Request,
	data interface{}, status int) {
	w.WriteHeader(status)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
