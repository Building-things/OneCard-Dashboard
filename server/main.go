package main

import (
	onecardparsing "onecard-dashboard/onecard-parsing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/", Root)

	app.Listen(":5000")
}

func Root(c *fiber.Ctx) error {
	payload := struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}{}

	err := c.BodyParser(&payload)
	if err != nil {
		return c.SendStatus(400)
	}
	if payload.Username == "test" {
		return c.JSON(onecardparsing.TestData())
	}

	data, ok := onecardparsing.OneCardData(payload.Username, payload.Password)
	if !ok {
		return c.SendStatus(400)
	}

	return c.JSON(data)
}
