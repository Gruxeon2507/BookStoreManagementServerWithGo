package bookHandler

import (
	"bookstore/ent"
	"bookstore/ent/book"
	"bookstore/ent/category"
	"bookstore/ent/user"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx, client *ent.Client) error {
	books, err := client.Book.Query().WithCategory().All(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "no books found",
			"data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "books found",
		"data":    books,
	})
}

func CreateBook(c *fiber.Ctx, client *ent.Client) error {
	type bookDTO struct {
		AuthorName  string
		CoverPath   string
		Description string
		Title       string
		PdfPath     string
		Price       float64
		Category    []string
		CreatedBy   string
	}
	var book bookDTO
	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "check your input",
			"error":   err,
		})
	}
	categories, err := client.Category.Query().Where(category.CategoryNameIn(book.Category...)).All(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error getting category",
			"error":   err,
		})
	}
	createdBy, err := client.User.Query().Where(user.Username(book.CreatedBy)).Only(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error getting user",
			"error":   err,
		})
	}
	b, err := client.Book.Create().
		SetAuthorName(book.AuthorName).
		SetCoverPath(book.CoverPath).
		SetDescription(book.Description).
		SetTitle(book.Title).
		SetPdfPath(book.PdfPath).
		AddCategory(categories...).
		SetPrice(book.Price).
		SetCreatedBy(createdBy.ID).
		Save(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "can not create new book",
			"error:":  err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "book created successfully",
		"data":    b,
	})

}
func GetBookById(c *fiber.Ctx, client *ent.Client) error {
	bookId, err := strconv.ParseInt(c.Params("bookId"), 10, 64)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "book not found",
			"error:":  err,
		})
	}
	book, err := client.Book.Query().Where(book.ID(int(bookId))).Only(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "query book error",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "book found",
		"data":    book,
	})

}

func UpdateBook(c *fiber.Ctx, client *ent.Client) error {

	bookId, err := strconv.ParseInt(c.Params("bookId"), 10, 64)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong id format",
			"error:":  err,
		})
	}
	bookFound, err := client.Book.Query().Where(book.ID(int(bookId))).Only(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error query book",
			"error":   err,
		})
	}
	if bookFound == nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "book not found",
			"error":   err,
		})
	}
	bookData := new(ent.Book)
	err = c.BodyParser(bookData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "missing/wrong book information",
			"error":   err,
		})
	}

	bookUpdate, err := client.Book.Update().
		Where(book.ID(int(bookId))).
		SetAuthorName(bookData.AuthorName).
		SetCoverPath(bookData.CoverPath).
		SetCreatedBy(bookData.CreatedBy).
		SetDescription(bookData.Description).
		SetTitle(bookData.Title).
		SetPdfPath(bookData.PdfPath).
		Save(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error updating book",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "book update successfully",
		"data":    bookUpdate,
	})
}

func DeleteBook(c *fiber.Ctx, client *ent.Client) error {
	bookId, err := strconv.ParseInt(c.Params("bookId"), 10, 32)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "wrong id format",
			"error:":  err,
		})
	}
	bookFound, err := client.Book.Query().Where(book.ID(int(bookId))).Only(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error query book",
			"error":   err,
		})
	}
	if bookFound == nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "book not found to delete",
			"error":   err,
		})
	}
	bookDeleted, err := client.Book.Delete().Where(book.ID(int(bookId))).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "error delete book",
			"error":   err,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "book deleted successfully",
		"data":    bookDeleted,
	})
}
