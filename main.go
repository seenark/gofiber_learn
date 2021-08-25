package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Get("/body", func(c *fiber.Ctx) error {

		person := map[string]interface{}{}
		err := c.BodyParser(&person)
		if err != nil {
			return err
		}

		return c.JSON(person)
	})

	app.Listen(":8080")
}

type Person struct {
	Name    string
	Surname string
}
