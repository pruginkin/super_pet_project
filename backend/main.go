package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{EnablePrintRoutes: true})

	app.Route("/", func(router fiber.Router) {
		router.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Привет мир")
		})

	})

	log.Fatal(app.Listen(":3000"))
}
