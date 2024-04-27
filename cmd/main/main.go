package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"simple_go_api/internal/handlers/actors"
	"simple_go_api/internal/handlers/interests"
	"simple_go_api/internal/handlers/languages"
	"simple_go_api/internal/handlers/tools"
	"time"
)

func main() {
	http.HandleFunc("/interests", interests.Handler)
	http.HandleFunc("/languages", languages.Handler)
	http.HandleFunc("/tools", tools.Handler)
	http.HandleFunc("/actors", actors.Handler)

	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server gracefully stopped")
}
