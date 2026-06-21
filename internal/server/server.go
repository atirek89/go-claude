// Package server provides a simple HTTP server for the go-claude project.
package server

import (
	"fmt"
	"net/http"

	"github.com/atirek89/go-claude/internal/greeting"
)

// New returns an http.ServeMux with all routes registered.
func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/goodbye", goodbyeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/", notFoundHandler)
	return mux
}

// Run starts the HTTP server on the given address (e.g. ":8080").
func Run(addr string, mux *http.ServeMux) error {
	fmt.Printf("Listening on http://localhost%s\n", addr)
	return http.ListenAndServe(addr, mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintln(w, greeting.Hello(name))
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintln(w, greeting.Goodbye(name))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status":"ok"}`)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 not found", http.StatusNotFound)
}
