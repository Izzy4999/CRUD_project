package middlewares

import (
	"github.com/Izzy499/crud_api/initializers"
	"github.com/gofiber/fiber/v2"
)

func LoggedIn() fiber.Handler {
	return checkLogin
}

func checkLogin(c *fiber.Ctx) error {
	sess, err := initializers.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	userId := sess.Get("userId")
	if userId == nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"error":   "cannot access login",
		})
	}

	c.Locals("userId", userId)
	return c.Next()
}
