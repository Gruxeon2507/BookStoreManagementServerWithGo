package userRoutes

import (
	"bookstore/ent"
	userHandler "bookstore/internal/handlers/user"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRotes(router fiber.Router, client *ent.Client) {
	user := router.Group("/user")

	//Register
	user.Post("/", func(c *fiber.Ctx) error {
		return userHandler.Register(c, client)
	})

	//Get All User
	user.Get("/", func(c *fiber.Ctx) error {
		return userHandler.GetAllUser(c, client)
	})

	//Get single User
	user.Get("/:userId", func(c *fiber.Ctx) error {
		return userHandler.GetUserByUsername(c, client)
	})

	//Update User
	user.Put("/:userId", func(c *fiber.Ctx) error {
		return userHandler.UpdateUser(c, client)
	})

	//Delete
	user.Delete("/:userId", func(c *fiber.Ctx) error {
		return userHandler.DeleteUser(c, client)
	})
}
