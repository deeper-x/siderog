package server

import (
	"io"
	"log"
	"net/http"
)

// StartSession creates the memory token
func StartSession() http.HandlerFunc {
	sFunc := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello session")
	}

	return sFunc
}

// RunServer is the http listener/server
func RunServer() {
	http.HandleFunc("/", StartSession())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
