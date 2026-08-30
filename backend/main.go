package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"bugsinguish-backend/db"
	"bugsinguish-backend/handlers"
	"bugsinguish-backend/sse"
)

func main() {
	ctx := context.Background()

	if err := db.Connect(ctx, os.Getenv("DATABASE_URL")); err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	r := chi.NewRouter()

	// Standard Chi middleware stack
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// CORS: allow any origin (local SvelteKit dev + Vercel production frontend)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Health check route (supports both GET and HEAD for Uptime monitors like UptimeRobot)
	r.Get("/health", healthCheckHandler)
	r.Head("/health", healthCheckHandler)

	// Ticket CRUD (dedup-check happens inside CreateTicket)
	r.Route("/tickets", func(r chi.Router) {
		r.Post("/", handlers.CreateTicket)
		r.Get("/", handlers.ListTickets)
		r.Get("/{id}", handlers.GetTicket)
		r.Patch("/{id}", handlers.UpdateTicket)
	})

	// Live progress log stream (sandbox + AI diagnosis events)
	r.Get("/api/stream", sse.StreamHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	log.Printf("bugsinguish backend listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"time":   time.Now().UTC().Format(time.RFC3339),
	})
}
