package main

import (
	onecardparsing "onecard-dashboard/onecard-parsing"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")

	app := fiber.New()

	app.Static("/", "./dist")

	app.Post("/", Root)

	app.Listen(":" + port)
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
