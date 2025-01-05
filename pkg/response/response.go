package response

import "github.com/gofiber/fiber/v2"

func JSON(c *fiber.Ctx, status int, data interface{}) error {
    return c.Status(status).JSON(fiber.Map{
        "status": status,
        "data":   data,
    })
}

func Error(c *fiber.Ctx, status int, message string) error {
    return c.Status(status).JSON(fiber.Map{
        "status":  status,
        "error":   true,
        "message": message,
    })
}