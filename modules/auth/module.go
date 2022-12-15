package auth

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	log.Println("REGISTER MODULE: Auth")

	api := app.Group("/api", AuthMiddlware)

	api.Get("/auth", func(c *fiber.Ctx) error {
		return c.SendString("AUTH")
	})
}
