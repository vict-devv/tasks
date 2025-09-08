// Package main is the entry point that leverages the API.
package main

import (
	"fmt"
	"os"

	"github.com/vict-devv/tasks/api"
	"github.com/vict-devv/tasks/internal/config"
)

func main() {
	cfg := config.New()
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	app := api.InitServer()
	if err := app.Listen(port); err != nil {
		// TODO: implement Logger and logging the Listen error
		os.Exit(1)
	}
}
