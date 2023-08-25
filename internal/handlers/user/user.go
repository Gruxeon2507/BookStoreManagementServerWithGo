package userHandler

import (
	"bookstore/ent"
	"bookstore/ent/user"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx, client *ent.Client) error {
	payload := new(ent.User)
	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "error parsing user data",
			"error":   err,
		})
	}
	user, err := client.User.Create().SetUsername(payload.Username).
		SetPassword(payload.Password).
		SetDisplayName(payload.DisplayName).
		SetDob(payload.Dob).
		SetEmail(payload.Email).
		SetCreatedDate(time.Now()).
		SetAvatarPath(payload.AvatarPath).
		Save(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error creating new user",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user created successfully",
		"data":    user,
	})
}

func GetAllUser(c *fiber.Ctx, client *ent.Client) error {
	users, err := client.User.Query().All(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error querying user",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "get user successful",
		"data":    users,
	})
}

func GetUserByUsername(c *fiber.Ctx, client *ent.Client) error {
	username := c.Params("userId")
	user, err := client.User.Query().WithCreatedBy().Where(user.Username(username)).Only(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error querying user",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "error querying user",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx, client *ent.Client) error {
	username := c.Params("userId")
	user, err := client.User.Query().Where(user.Username(username)).Only(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error querying user",
			"error":   err,
		})
	}
	if user == nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "user not exist",
		})
	}
	payload := new(ent.User)
	err = c.BodyParser(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "user input data is not valid",
			"error":   err,
		})
	}
	updatedUser, err := client.User.Update().
		SetUsername(payload.Username).
		SetPassword(payload.Password).
		SetDisplayName(payload.DisplayName).
		SetDob(payload.Dob).
		SetEmail(payload.Email).
		SetCreatedDate(time.Now()).
		SetAvatarPath(payload.AvatarPath).
		Save(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "update user failed",
			"err":     err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user updated successfully",
		"data":    updatedUser,
	})
}

func DeleteUser(c *fiber.Ctx, client *ent.Client) error {
	username := c.Params("userId")
	userFind, err := client.User.Query().Where(user.Username(username)).Only(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error querying user",
			"error":   err,
		})
	}
	if userFind == nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "user not exist",
		})
	}
	deletedUser, err := client.User.Delete().Where(user.Username(username)).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error deleting user",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user deleted successfully",
		"data":    deletedUser,
	})
}
