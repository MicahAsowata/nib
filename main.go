package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello 🌍🌍🌍")
	})

	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		log.Panic(err)
	}
}
