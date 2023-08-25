package router

import (
	"bookstore/ent"
	bookRoutes "bookstore/internal/routes/book"
	categoryRoutes "bookstore/internal/routes/category"
	userRoutes "bookstore/internal/routes/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, client *ent.Client) {
	api := app.Group("/api", logger.New())
	categoryRoutes.SetupCategoryRoutes(api, client)
	bookRoutes.SetUpBookRoutes(api, client)
	userRoutes.SetUpUserRotes(api, client)
}
