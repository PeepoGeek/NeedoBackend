package main

import (
	"log"
	"net/http"

	_ "github.com/example/myproject/docs" // This will be auto-generated
	"github.com/example/myproject/internal/healthcheck"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title NeedoBackend API
// @version 1.0
// @description This is the API documentation for the NeedoBackend service
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()

	// CORS middleware configuration
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"}),
	)

	// Apply CORS middleware to router
	r.Use(corsMiddleware)

	// API Routes
	r.HandleFunc("/healthcheck", healthcheck.Handler).Methods("GET")

	// Swagger documentation - updated configuration
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	log.Println("Server running on http://localhost:8080")
	log.Println("Swagger documentation available at http://localhost:8080/swagger/index.html")

	// Wrap the router with CORS middleware
	handler := corsMiddleware(r)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
