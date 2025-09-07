// Package main is the entry point that leverages the API.
package main

import (
	"os"

	"github.com/vict-devv/tasks/api"
)

func main() {
	app := api.InitServer()
	if err := app.Listen(":3000"); err != nil {
		os.Exit(1)
	}
}
