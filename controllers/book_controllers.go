package controllers

import (
	"net/http"
	"playgofiber/config"
	"playgofiber/entity"
	"playgofiber/helpers"
	"playgofiber/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.JSON("welcome to api")
}

func List(c *fiber.Ctx) error {

	var books []models.Book

	config.InitDB().Debug().Scopes(helpers.Paginate(c)).Find(&books)

	return c.JSON(helpers.ResponseSuccess("sucess",fiber.StatusOK,books))
}

func ShowBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var books models.Book

	if id == "" {
		return c.JSON(helpers.ResponseFailed("error",fiber.StatusUnprocessableEntity,"invalid id"))
	}

	config.InitDB().First(&books,id)

	if books.BookName == "" {
		return c.JSON(helpers.ResponseFailed("error",fiber.StatusBadRequest,"data not found"))
	}

	return c.JSON(helpers.ResponseSuccess("success",fiber.StatusOK,books))
}

func CreateBook(c *fiber.Ctx) error {
	var books models.Book

	p := new(entity.CreateBookRequest)

	if err := c.BodyParser(p); err != nil {
		return c.JSON(helpers.ResponseFailed("error",fiber.StatusInternalServerError,err.Error()))
	}

	config.InitDB().Create(&models.Book{
		BookName:p.BookName,
		BookAuthor: p.BookAuthor, 
		BookYear: p.BookYear})

	errors := config.ValidateStruct(*p)
    if errors != nil {
        return c.JSON(helpers.ResponseFailed("error",fiber.StatusBadRequest,errors))
    }

	config.InitDB().Where("book_name = ?",p.BookName).First(&books)

	response := helpers.ResponseSuccess("success",fiber.StatusOK,books)

	return c.JSON(response)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	p := new(entity.UpdateBookRequest)

	if err := c.BodyParser(p); err != nil {
		return c.JSON(helpers.ResponseFailed("error",fiber.StatusInternalServerError,err.Error()))
	}
	
	if id == "" {
		return c.JSON(helpers.ResponseFailed("error",fiber.StatusUnprocessableEntity,"invalid id"))
	}

	config.InitDB().Model(&models.Book{}).Where("id = ?",id).Updates(&models.Book{
		BookName: p.BookName,
		BookAuthor: p.BookAuthor,
		BookYear: p.BookYear})

	var books models.Book

	config.InitDB().First(&books,id)

	response := helpers.ResponseSuccess("success",fiber.StatusOK,books)

	return c.JSON(response)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.JSON(helpers.ResponseFailed("error",fiber.StatusUnprocessableEntity,"invalid id"))
	}

	config.InitDB().Delete(&models.Book{},id)

	return c.JSON(&fiber.Map{
		"status": "success",
		"code": http.StatusOK,
		"message": "Data has been deleted"})
}