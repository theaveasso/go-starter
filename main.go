package main

import (
	"log/slog"
	"os"

	"github.com/theaveasso/go-starter/config"
	"github.com/theaveasso/go-starter/internal/application"
)

func main() {
  config, err := config.LoadConfig(".")
  if err != nil {
    slog.Error("fail to load config", "error", err)
    os.Exit(1)
  }


	app := application.NewApplication(config)

	if err := app.Start(); err != nil {
		slog.Error("fail to start server", "error", err)
	}
}
