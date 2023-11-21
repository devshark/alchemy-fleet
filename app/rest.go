package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devshark/alchemy-fleet/adapters"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HttpServer struct {
	spacecraftService adapters.Spacecraft
}

const maxHeaderBytes = 1 << 20

func NewHTTPServer(spacecraftService adapters.Spacecraft, port int, defaultTimeout time.Duration) *http.Server {
	httpServer := http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           router(spacecraftService),
		ReadTimeout:       defaultTimeout,
		WriteTimeout:      defaultTimeout,
		ReadHeaderTimeout: defaultTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	return &httpServer
}

func router(spacecraftService adapters.Spacecraft) *chi.Mux {
	mux := chi.NewRouter()

	server := &HttpServer{
		spacecraftService: spacecraftService,
	}

	mux.Use(middleware.Logger)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Heartbeat("/ping"))

	// public endpoints
	mux.Group(func(mux chi.Router) {
		mux.Get("/spacecraft/{id}", server.GetSpacecraftById)
		mux.Get("/spacecraft", server.ListSpacecrafts)
	})

	// auth endpoints
	mux.Group(func(mux chi.Router) {
		// credentials are hardcoded for simplicity
		creds := map[string]string{
			"username": "password",
		}
		mux.Use(middleware.BasicAuth("spacecraft", creds))
		mux.Post("/spacecraft", server.CreateSpacecraft)
		mux.Put("/spacecraft/{id}", server.UpdateSpacecraft)
		mux.Delete("/spacecraft/{id}", server.DeleteSpacecraft)
	})

	return mux
}

func respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type errorResponse struct {
	Error string `json:"error"`
}

type genericResponse struct {
	Success bool `json:"success"`
}
