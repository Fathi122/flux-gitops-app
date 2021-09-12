package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// hello
func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome Page !"}`))
}
func router() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/", hello)
	return m
}
func main() {
	const addr = ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: router(),
	}
	log.Infof("Listening on '%s'", addr)
	log.Fatal(server.ListenAndServe())
}
