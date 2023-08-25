package categoryHandler

import (
	"bookstore/ent"
	"bookstore/ent/category"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx, client *ent.Client) error {
	categoryData := new(ent.Category)
	err := c.BodyParser(categoryData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"error":   err,
		})
	}
	category, err := client.Category.Create().
		SetCategoryName(categoryData.CategoryName).
		Save(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error creating category",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "category created successfully",
		"data":    category,
	})
}

func GetAllCategory(c *fiber.Ctx, client *ent.Client) error {
	categories, err := client.Category.Query().All(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "error getting category",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "error",
		"message": "get category success",
		"data":    categories,
	})
}
func GetAllBookCategory(c *fiber.Ctx, client *ent.Client) error {
	categoryName := c.Params("categoryName")
	categoryFound, err := client.Category.Query().Where(category.CategoryName(categoryName)).Only(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "error getting category",
			"error":   err,
		})
	}
	if categoryFound == nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "success",
			"message": "category not found",
			"error":   err,
		})
	}
	books, err := client.Category.QueryBook(categoryFound).All(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "query book error",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "book found",
		"data":    books,
	})
}

func UpdateCategory(c *fiber.Ctx, client *ent.Client) error {
	categoryName := c.Params("categoryName")
	categoryFound, err := client.Category.Query().Where(category.CategoryName(categoryName)).Only(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "error getting category",
			"error":   err,
		})
	}
	if categoryFound == nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "category not found",
			"error":   err,
		})
	}
	updateData := new(ent.Category)
	err = c.BodyParser(updateData)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "category not found",
			"error":   err,
		})
	}
	updatedCategory, err := client.Category.Update().SetCategoryName(updateData.CategoryName).Save(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "update category error",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "update category successfully",
		"data":    updatedCategory,
	})
}

func DeleteCategory(c *fiber.Ctx, client *ent.Client) error {
	categoryName := c.Params("categoryName")
	categoryFound, err := client.Category.Query().Where(category.CategoryName(categoryName)).Only(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "error getting category",
			"error":   err,
		})
	}
	if categoryFound == nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "category not found",
			"error":   err,
		})
	}
	deletedCategory, err := client.Category.Delete().Where(category.CategoryName(categoryName)).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error delete category",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "delete category successfully",
		"data":    deletedCategory,
	})
}
