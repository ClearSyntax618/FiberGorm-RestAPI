package main

import (
	"log"

	"github.com/ClearSyntax618/FiberGorm-RestAPI/config"
	"github.com/ClearSyntax618/FiberGorm-RestAPI/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/", handlers.HelloWorld)
	app.Post("/dog", handlers.CreateDog)
	app.Get("/dogs", handlers.GetDogs)
	app.Get("/get/:id", handlers.GetDog)
	app.Put("/update/:id", handlers.UpdateDog)
	app.Delete("/delete/:id", handlers.DeleteDog)

	log.Fatal(app.Listen(":3000"))
}
