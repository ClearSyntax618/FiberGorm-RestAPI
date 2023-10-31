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

func GetDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog models.Dog

	result := config.DB.Find(&dog, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"msg": "Not found",
		})
	}

	return c.Status(200).JSON(&dog)
}

func UpdateDog(c *fiber.Ctx) error {
	dog := new(models.Dog)
	id := c.Params("id")

	if err := c.BodyParser(dog); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": err,
		})
	}

	config.DB.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func DeleteDog(c *fiber.Ctx) error {
	id := c.Params("id")
	var dog models.Dog

	result := config.DB.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"msg": "Not Found",
		})
	}

	return c.SendStatus(200)
}
