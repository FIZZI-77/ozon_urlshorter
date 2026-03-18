package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"ozon/pkg"
	"ozon/src/core/handler"
	"ozon/src/core/repository"
	"ozon/src/core/service"
	"sync"
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
	defer db.Close()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handl := handler.NewHandler(services)

	srv := new(Server)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := srv.run(os.Getenv("SERVER_PORT"), handl.InitRoutes()); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error starting server: %v", err)
		}

	}()

	log.Printf("Server started on port %s", os.Getenv("PORT_SERVER"))

	wg.Wait()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = srv.shutdown(shutdownCtx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}
	if err = db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

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
