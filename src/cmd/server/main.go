package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"ozon/pkg"
	"ozon/src/core/handler"
	"ozon/src/core/repository"
	"ozon/src/core/service"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

var StorageType = flag.String("n", "postgres", "choose storage postgres/memory")

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	flag.Parse()
	var repo *repository.Repository

	switch *StorageType {
	case "postgres":
		db, err := pkg.NewPostgresDB(pkg.Config{
			Host:     os.Getenv("HOST"),
			Port:     os.Getenv("PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("SSLMODE"),
		})
		defer func() {
			if err := db.Close(); err != nil {
				log.Printf("Error closing database: %v", err)
			}
		}()
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		repo = repository.NewPostgresRepository(db)
	case "memory":
		repo = repository.NewMemoryRepository()
		fmt.Println("using memory storage")
	}

	services := service.NewService(repo)
	handl := handler.NewHandler(services)

	srv := new(Server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.run(os.Getenv("SERVER_PORT"), handl.InitRoutes()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}

	}()

	log.Printf("Server started on port %s", os.Getenv("SERVER_PORT"))

	<-quit
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	log.Println("Shutting down server...")

	if err := srv.shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}

	log.Println("Server stopped")
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
