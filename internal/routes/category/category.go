package categoryRoutes

import (
	"bookstore/ent"
	categoryHandler "bookstore/internal/handlers/category"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(router fiber.Router, client *ent.Client) {
	category := router.Group("/category")

	//Create category
	category.Post("/", func(c *fiber.Ctx) error {
		return categoryHandler.CreateCategory(c, client)
	})

	//Get all category
	category.Get("/", func(c *fiber.Ctx) error {
		return categoryHandler.GetAllCategory(c, client)
	})

	//Get all book in a category
	category.Get("/:categoryName", func(c *fiber.Ctx) error {
		return categoryHandler.GetAllBookCategory(c, client)
	})

	//Update a category
	category.Put("/:categoryName", func(c *fiber.Ctx) error {
		return categoryHandler.UpdateCategory(c, client)
	})

	//Delete one category
	category.Delete("/:categoryName", func(c *fiber.Ctx) error {
		return categoryHandler.DeleteCategory(c, client)
	})
}
