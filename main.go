package main

import (
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")
}
