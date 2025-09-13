// Package main is the entry point that leverages the API.
package main

import (
	"fmt"

	"github.com/vict-devv/tasks/api"
	"github.com/vict-devv/tasks/internal/config"
	"github.com/vict-devv/tasks/internal/constants"
	"github.com/vict-devv/tasks/internal/logger"
)

func main() {
	cfg := config.New()
	logger := logger.NewStandardLogger()

	addr := fmt.Sprintf("%s:%d", cfg.Database.Host, cfg.Server.Port)
	app := api.InitServer()

	logger.Info(constants.LogPackageMain, "Starting server")
	if err := app.Listen(addr); err != nil {
		logger.Fatalf(constants.LogPackageMain, "Server has stopped with error: %w", err)
	}
}

// TODO:
// - [ ] rename tasks folder in cmd to tasks-api
// - [ ] create Dockerfile, Dockerignore and Docker-compose files
// - [ ] implement graceful shutdown
// - [ ] implement more logging
// - [ ] implement tests
// - [ ] implement CI/CD
// - [ ] implement more features
