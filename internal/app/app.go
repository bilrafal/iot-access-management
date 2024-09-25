package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	config2 "iot-access-management/internal/config"
	"net/http"
	"time"
)

type App struct {
	router *chi.Mux
	config config2.Config
}

func New(configPath string, routeDefs []RouteDef) *App {
	config := config2.LoadConfig(configPath)
	router := LoadGroupOfRoutes(routeDefs)
	app := &App{
		config: config,
		router: router,
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	var err error

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.ServerPort),
		Handler: a.router,
	}

	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(),
			time.Second*time.Duration(a.config.Server.Timeout))
		defer cancel()

		return server.Shutdown(timeout)
	}
}
