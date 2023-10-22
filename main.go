package main

import (
	"github.com/ClearSyntax618/FiberGorm-RestAPI/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handlers.HelloWorld)

	app.Listen(":3000")
}
