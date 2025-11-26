package http

import (
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TrxController struct {
	usecase *usecase.TrxUsecase
	log     *logrus.Logger
}

func NewTrxController(usecase *usecase.TrxUsecase, log *logrus.Logger) *TrxController {
	return &TrxController{usecase, log}
}

func (c *TrxController) Create(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)

	var request model.CreateTransactionRequest
	if err := ctx.BodyParser(&request); err != nil {
		c.log.Warnf("Invalid checkout body: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := c.usecase.Checkout(userId, request)
	if err != nil {
		c.log.Errorf("Checkout failed: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   response,
	})
}

func (c *TrxController) History(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)

	responses, err := c.usecase.History(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"data":   responses,
	})
}
