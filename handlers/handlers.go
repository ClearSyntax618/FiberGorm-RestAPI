package handlers

import (
	"github.com/ClearSyntax618/FiberGorm-RestAPI/config"
	"github.com/ClearSyntax618/FiberGorm-RestAPI/models"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hola, mundillo!")
}

func CreateDog(c *fiber.Ctx) error {
	dog := new(models.Dog)

	if err := c.BodyParser(dog); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	config.DB.Create(&dog)
	return c.Status(200).JSON(dog)
}

func GetDogs(c *fiber.Ctx) error {
	var dogs []models.Dog
	var err error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	config.DB.Find(&dogs)
	return c.Status(200).JSON(dogs)
}
