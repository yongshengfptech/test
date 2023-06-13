package router

import (
	"api_test/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup_router(app *fiber.App) {
	staff := app.Group("/staff")
	staff.Post("/add", handler.CreateStaff)
}
