package http

import (
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AddressController struct {
	usecase *usecase.AddressUsecase
	log     *logrus.Logger
}

func NewAddressController(usecase *usecase.AddressUsecase, log *logrus.Logger) *AddressController {
	return &AddressController{usecase: usecase, log: log}
}

func (c *AddressController) List(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)

	responses, err := c.usecase.List(userId)
	if err != nil {
		c.log.Warnf("Failed to list addresses: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   responses,
	})
}

func (c *AddressController) Create(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)
	var request model.CreateAddressRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := c.usecase.Create(userId, request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": response})
}

func (c *AddressController) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(uint)
	id, _ := ctx.ParamsInt("id")

	err := c.usecase.Delete(userId, uint(id))
	if err != nil {
		if err.Error() == "forbidden: anda tidak berhak menghapus alamat ini" {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(fiber.Map{"data": "OK"})
}
