package handler

import (
	"api_test/database"
	"api_test/logger"
	"api_test/model"
	"api_test/response"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IncomingStaff struct {
	Staff_name string `json:"staff_name"`
}

func CreateStaff(ctx *fiber.Ctx) error {
	db := database.DB

	var request_body *IncomingStaff
	if err := ctx.BodyParser(&request_body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Json_msg("error", "fail to get request"))
	}

	staff_id := uuid.New()

	staff_data := model.Staff{
		ID:          staff_id,
		Create_time: time.Now(),
		Staff_name:  request_body.Staff_name,
	}

	err := db.Create(staff_data).Error
	if err != nil {
		logger.Log.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Json_msg("error", "fail to create data"))
	}

	logger.Log.Info("create succeed")
	return ctx.Status(fiber.StatusOK).JSON(response.Json_msg("success", "create succeed "+request_body.Staff_name))
}
