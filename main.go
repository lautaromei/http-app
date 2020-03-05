package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func run() error {
	srv := newAboutServer()
	log.Fatal(http.ListenAndServe(":8080", srv))

	return nil
}

func newAboutServer() *about.server {
	s := &about.server{}
	s.router = mux.NewRouter()

	s.routes()
	return s
}
