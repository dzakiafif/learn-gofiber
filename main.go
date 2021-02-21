package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "playgofiber/routes"
)

func main() {
    app := fiber.New()

    app.Use(logger.New())

    routes.Routes(app)

    app.Listen(":3000")
}