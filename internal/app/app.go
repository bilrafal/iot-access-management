package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"iot-access-management/internal/binder"
	"iot-access-management/internal/config"
	"iot-access-management/internal/router"
	"net/http"
	"time"
)

type App struct {
	ctx    context.Context
	binder binder.Binder
	router *chi.Mux
	config config.Config
}

func New(ctx context.Context, appBinder binder.Binder, routeDefs []router.RouteDef) *App {
	cfg := appBinder.GetConfig()
	boundRoutes := appBinder.BindDependencies(ctx, routeDefs)
	appRouter := router.LoadGroupOfRoutes(boundRoutes)
	app := &App{
		ctx:    ctx,
		binder: appBinder,
		config: cfg,
		router: appRouter,
	}

	return app
}

func (a *App) Start() error {
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
	case <-a.ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(),
			time.Second*time.Duration(a.config.Server.Timeout))
		defer cancel()

		return server.Shutdown(timeout)
	}
}
