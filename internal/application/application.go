package application

import (
	"log/slog"
	"net/http"

	"github.com/theaveasso/go-starter/config"
)

type Application struct {
  config *config.Config
	router http.Handler
}

func NewApplication(config *config.Config) *Application {
	return &Application{
    config: config,
  }
}

func (app *Application) Start() error {
	server := http.Server{
		Addr:    app.config.Address,
		Handler: setupRoutes(),
	}

	slog.Info("server listening on", "port", app.config.Address)
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
