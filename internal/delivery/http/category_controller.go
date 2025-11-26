package http

import (
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryController struct {
	usecase *usecase.CategoryUsecase
	log     *logrus.Logger
}

func NewCategoryController(usecase *usecase.CategoryUsecase, log *logrus.Logger) *CategoryController {
	return &CategoryController{usecase, log}
}

func (c *CategoryController) Create(ctx *fiber.Ctx) error {
	var request model.CreateCategoryRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := c.usecase.Create(request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": response})
}

func (c *CategoryController) List(ctx *fiber.Ctx) error {
	responses, err := c.usecase.List()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(fiber.Map{"data": responses})
}

func (c *CategoryController) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	var request model.UpdateCategoryRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := c.usecase.Update(uint(id), request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": response,
	})
}

func (c *CategoryController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid category ID")
	}

	err = c.usecase.Delete(uint(id))
	if err != nil {
		c.log.Errorf("Failed to delete category: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": true,
	})
}
