package routes

import (
	"github.com/gofiber/fiber/v2"
	"playgofiber/controllers"
)

func Routes(app *fiber.App){

	api := app.Group("v1")

	api.Get("/",controllers.Index)

	api.Get("/list-book",controllers.List)

	api.Post("/create-book",controllers.CreateBook)

	api.Put("update-book/:id",controllers.UpdateBook)

	api.Delete("/delete-book/:id",controllers.DeleteBook)
}