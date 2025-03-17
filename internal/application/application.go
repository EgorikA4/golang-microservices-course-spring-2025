package application

import (
	"golang-microservices-course-spring-2025/internal/server/handlers"

	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()

	app.Get("/api/v1/contact", handlers.GetContact)
	app.Post("/api/v1/contact", handlers.CreateContact)
	app.Put("/api/v1/contact", handlers.UpdateContact)
	app.Delete("/api/v1/contact", handlers.DeleteContact)

	app.Get("/api/v1/group", handlers.GetGroup)
	app.Post("/api/v1/group", handlers.CreateGroup)
	app.Put("/api/v1/group", handlers.UpdateGroup)
	app.Delete("/api/v1/group", handlers.DeleteGroup)
	return app
}
