package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/B3ns44d/check-location/utils"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ping Pong!")
	})

	app.Get("/geo/:ip/:fields?", utils.Cache(10*time.Minute), utils.GetGeo())

	log.Fatal(app.Listen(":3000"))
}
