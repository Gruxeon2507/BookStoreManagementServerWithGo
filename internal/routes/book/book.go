package bookRoutes

import (
	"bookstore/ent"
	bookHandler "bookstore/internal/handlers/book"

	"github.com/gofiber/fiber/v2"
)

func SetUpBookRoutes(router fiber.Router, client *ent.Client) {
	book := router.Group("/book")

	//Create a Book
	book.Post("/", func(c *fiber.Ctx) error {
		return bookHandler.CreateBook(c, client)
	})

	//Get all books
	book.Get("/", func(c *fiber.Ctx) error {
		return bookHandler.GetBooks(c, client)
	})

	//Get a book
	book.Get("/:bookId", func(c *fiber.Ctx) error {
		return bookHandler.GetBookById(c, client)
	})

	//Update a book
	book.Post("/:bookId", func(c *fiber.Ctx) error {
		return bookHandler.UpdateBook(c, client)
	})

	//Delete a book
	book.Delete("/:bookId", func(c *fiber.Ctx) error {
		return bookHandler.DeleteBook(c, client)
	})
}
