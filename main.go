package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
}

func run() error {
	srv := about.newServer()
	log.Fatal(http.ListenAndServe(":8080", srv))

	return nil
}
