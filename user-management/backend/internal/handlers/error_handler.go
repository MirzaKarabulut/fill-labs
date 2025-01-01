package handlers

import "github.com/gofiber/fiber/v2"

func errorResponse(err error) fiber.Map {
	return fiber.Map{"error": err.Error()}
}
