package routes

import (
	"github.com/MirzaKarabulut/fill-labs/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", handlers.ListUsers)
	api.Get("/users/:id", handlers.GetUserById)
	api.Post("/users", handlers.CreateUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.DeleteUser)
}
