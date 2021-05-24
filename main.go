package main

import (
	"go-fiber/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.CustomerRouter(app)

	log.Fatal(app.Listen(":3000"))
}
