package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	const PORT = ":42069"
	app := fiber.New()

	log.Fatal(app.Listen(PORT))
}
