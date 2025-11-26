package middleware

import (
	"API-Ecommerce-Evermos/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(userRepo repository.IUserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIdVal := c.Locals("user_id")
		if userIdVal == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": "Unauthorized"})
		}
		userId := userIdVal.(uint)

		user, err := userRepo.FindById(userId)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"errors": "User not found"})
		}

		if !user.IsAdmin {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"errors": "Forbidden: Admin access required"})
		}
		
		return c.Next()
	}
}
