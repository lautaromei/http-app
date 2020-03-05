package about

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func (s *server) routes() {
	s.router.HandleFunc("/about", s.handleAbout("Hi!"))
}

func (s *server) handleAbout(format string) http.HandlerFunc {
	type response struct {
		Greeting string `json:"greeting"`
	}
	resp := &response{Greeting: format}

	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, resp, 200)
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
