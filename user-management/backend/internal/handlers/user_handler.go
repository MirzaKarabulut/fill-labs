package handlers

import (
	"github.com/MirzaKarabulut/fill-labs/internal/db"
	"github.com/gofiber/fiber/v2"
)

func ListUsers(ctx *fiber.Ctx) error {
	user := []db.User{}

	if err := db.DB.Find(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := db.User{}

	if err := db.DB.First(&user, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(ctx *fiber.Ctx) error {
	user := db.User{}

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := db.User{}

	if err := db.DB.First(&user, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	updateData := db.User{}
	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
	}

	if err := db.DB.Model(&user).Updates(updateData).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	updatedUser := db.User{}
	if err := db.DB.First(&updatedUser, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedUser)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := db.User{}

	if err := db.DB.First(&user, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
