package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	const PORT = ":42069"

	log.Fatal(app.Listen(PORT))
}
