package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleAbout(t *testing.T) {
	srv := newServer()

	r := httptest.NewRequest("GET", "/about", nil)
	w := httptest.NewRecorder()

	srv.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Fatal()
	}
}
