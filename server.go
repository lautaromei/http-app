package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func (s *server) routes() {
	s.router.HandleFunc("/test", s.handleTest("test"))
}

func run() error {
	srv := newServer()

	log.Fatal(http.ListenAndServe(":8080", srv))

	return nil
}

func newServer() *server {
	s := &server{}
	s.router = mux.NewRouter()

	s.routes()
	return s
}

func (s *server) handleTest(format string) http.HandlerFunc {
	type response struct {
		Greeting string `json:"greeting"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
