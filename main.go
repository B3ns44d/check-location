package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/B3ns44d/check-location/utils"
)

const idleTimeout = 5 * time.Second

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ping Pong!")
	})

	app.Get("/geo/:ip/:fields?", utils.Cache(10*time.Minute), utils.GetGeo())

	file, err := os.OpenFile("./logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	go func() {
		if err := app.Listen(":3005"); err != nil {
			log.Panic(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	file.Close()

	fmt.Println("App was successful shutdown.")
}
