package main

import (
	"github.com/ClearSyntax618/FiberGorm-RestAPI/config"
	"github.com/ClearSyntax618/FiberGorm-RestAPI/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/", handlers.HelloWorld)
	app.Post("/add-dog", handlers.CreateDog)
	app.Get("/get-dog", handlers.GetDogs)

	app.Listen(":3000")
}
