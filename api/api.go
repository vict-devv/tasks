// Package api provides all API related functionality.
package api

import "github.com/gofiber/fiber/v2"

// InitServer initializes and returns a Fiber app instance.
func InitServer() *fiber.App {
	server := fiber.New()

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return server
}
