package response

import (
	"github.com/gofiber/fiber/v2"
)

func Json_msg(status string, msg string) *fiber.Map {
	return &fiber.Map{
		"status": status,
		"msg":    msg,
	}
}
