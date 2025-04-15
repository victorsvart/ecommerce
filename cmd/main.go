package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/victorsvart/go-ecommerce/database"
	"github.com/victorsvart/go-ecommerce/wiring"
)

func setupApiBase() chi.Router {
	chi := chi.NewRouter()
	chi.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(60*time.Second),
	)

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
	db := database.Connect()
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
