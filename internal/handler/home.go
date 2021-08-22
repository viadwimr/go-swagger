
package handler

import "github.com/gofiber/fiber/v2"

func HomeHandler(f *fiber.App) {
	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}