package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/victorsvart/egommerce/internal/adapter/postgres"
	"github.com/victorsvart/egommerce/internal/wiring"
)

func setupApiBase() chi.Router {
	chi := chi.NewRouter()
	chi.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
		middleware.Timeout(60*time.Second),
	)

	chi.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	return chi
}

func logChiRoutes(r chi.Router) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("Routed %s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}

func main() {
	db := postgres.Connect()
	api := setupApiBase()
	api.Route("/v1/api", func(r chi.Router) {
		wiring.WireApp(db, r)
	})
	logChiRoutes(api)

	server := http.Server{
		Addr:    ":8080",
		Handler: api,
	}
	log.Println("Server started at :8080")
	server.ListenAndServe()
}
