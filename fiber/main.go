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

	// description {write} post request
	app.Post("/write", func(c *fiber.Ctx) error {
		payload := new(utils.CommentDoc)
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":     "failure",
				"error":      "Invalid request body",
				"statusCode": fiber.StatusBadRequest,
			})
		}
		return c.JSON(payload)
	})
	err := app.Listen(port)
	if err != nil {
		panic(err)
	}
}
