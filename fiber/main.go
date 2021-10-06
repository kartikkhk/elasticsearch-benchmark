package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kartikkhk/elasticsearch-benchmarks/utils"
)

const port = ":3000"

func main() {
	app := fiber.New()
	app.Use(logger.New())
	ctx := context.Background()
	esclient, err := utils.GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic(err)
	}

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("pong")
	})

	// description {write} post request
	app.Post("/write", func(c *fiber.Ctx) error {
		payload := new(utils.CommentDoc)
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":     "failure",
				"error":      "invalid request body",
				"statusCode": fiber.StatusBadRequest,
			})
		}
		dataJSON, _ := json.Marshal(*payload)
		js := string(dataJSON)
		reply, err := esclient.Index().
			Index("comments").
			BodyJson(js).
			Type("comment").
			Do(ctx)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":     "failure",
				"error":      "internal server error",
				"statusCode": fiber.StatusInternalServerError,
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":     "success",
			"statusCode": fiber.StatusCreated,
			"result":     reply,
		})

	})

	app.Listen(port)
}
