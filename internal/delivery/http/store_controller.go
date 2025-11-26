package http

import (
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type StoreController struct {
	usecase *usecase.StoreUsecase
	log     *logrus.Logger
}

func NewStoreController(usecase *usecase.StoreUsecase, log *logrus.Logger) *StoreController {
	return &StoreController{usecase, log}
}

func (c *StoreController) GetMyStore(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)

	response, err := c.usecase.GetMyStore(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.JSON(fiber.Map{"data": response})
}

func (c *StoreController) Update(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)

	var request model.UpdateStoreRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := c.usecase.Update(userId, request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{"data": response})
}
