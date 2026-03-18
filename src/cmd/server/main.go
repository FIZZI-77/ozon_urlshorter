package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ozon/pkg"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := pkg.NewPostgresDB(pkg.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Received termination signal, shutting down...")
}

func (s *Server) run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
