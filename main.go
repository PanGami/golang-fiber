package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/api/check", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Golang, Fiber, Gorm Go Project has been successfully",
		})
	})
	log.Fatal(app.Listen("localhost:8000"))
}
