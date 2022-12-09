package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"fiberapp/modules/auth"
	"fiberapp/modules/public"
)

func main() {
	app := fiber.New()

	firebaseAuth := auth.Register(app)

	log.Println(firebaseAuth)

	app.Use(limiter.New(limiter.Config{
		Expiration: 1 * time.Hour,
		Max:        200,
	}))
	app.Use(requestid.New())

	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	public.Register(app)

	log.Fatal(app.Listen(":3000"))
}
