package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Veylin2175/alotofmodules/http-server/internal/config"
)

type Server struct {
	ctx    context.Context
	client *http.Server
}

func main() {
	// TODO: Нужно перенести main в другую директории
	// TODO: Также добавить прогрузку локального конфига

	ctx := context.Background()
	cfg := &config.Config{
		LogLevel: "DEBUG",
		HttpServer: config.HttpServer{
			Host: "localhost",
			Port: 8596,
		},
	}

	server := NewServer(ctx, cfg)

	err := server.RunServer(ctx)
	if err != nil {
		log.Fatalf("server run error: %v", err)
	}
}

func NewServer(ctx context.Context, cfg *config.Config) *Server {
	const fn = "NewServer"

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello from my server"))
		if err != nil {
			fmt.Printf("%s : %s", fn, err.Error())
			return
		}
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: router,
	}

	return &Server{
		ctx:    ctx,
		client: server,
	}
}

func (s *Server) RunServer(ctx context.Context) error {
	const fn = "RunServer"

	err := s.client.ListenAndServe()
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	return nil
}
