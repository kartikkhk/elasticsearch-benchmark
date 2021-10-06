package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kartikkhk/elasticsearch-benchmarks/utils"
)

const port = ":3000"

func main() {
	app := fiber.New()
	app.Use(logger.New())
	utils.ConstructClient()
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("pong")
	})
	err := app.Listen(port)
	if err != nil {
		panic(err)
	}
}
