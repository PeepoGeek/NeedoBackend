package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/example/myproject/internal/healthcheck"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", healthcheck.Handler).Methods("GET")

	log.Println("Server corriendo en el puerto :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
} 