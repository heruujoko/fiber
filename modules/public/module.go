package public

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	log.Println("REGISTER MODULE: Public")
	app.Get("/", func(c *fiber.Ctx) error {
		resp := ApiVersion{
			Name:    "FiberApp",
			Version: "0.1.0",
		}
		return c.JSON(resp)
	})
}
