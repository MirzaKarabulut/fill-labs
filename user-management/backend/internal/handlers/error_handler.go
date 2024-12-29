package handlers

import "github.com/gofiber/fiber"

func errorResponse(err error) fiber.Map {
	return fiber.Map{"error": err.Error()}
}
