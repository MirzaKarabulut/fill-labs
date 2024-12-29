package handlers

import (
	"github.com/MirzaKarabulut/fill-labs/internal/db"
	"github.com/gofiber/fiber"
)

func ListUsers(ctx *fiber.Ctx) {
	user := []db.User{}

	if err := db.DB.Find(&user).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	ctx.Status(fiber.StatusOK).JSON(user)
}

func GetUserById(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	user := db.User{}

	if err := db.DB.First(&user, id).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	ctx.Status(fiber.StatusOK).JSON(user)
}

func CreateUser(ctx *fiber.Ctx) {
	user := db.User{}

	if err := ctx.BodyParser(&user); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	ctx.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	user := db.User{}

	if err := db.DB.First(&user, id).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	updateData := db.User{}
	if err := ctx.BodyParser(&updateData); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(errorResponse(err))
		return
	}

	if err := db.DB.Model(&user).Updates(updateData).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	updatedUser := db.User{}
	if err := db.DB.First(&updatedUser, id).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	ctx.Status(fiber.StatusOK).JSON(updatedUser)
}

func DeleteUser(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	user := db.User{}

	if err := db.DB.First(&user, id).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
		return
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
