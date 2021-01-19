package main

import (
	"github.com/Webservice-Rathje/Mailing-Service/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", controller.Default)
	app.Get("/sendRegistrationCodeMail", controller.SendRegistrationCodeController)

	app.Listen(":8081")
}
