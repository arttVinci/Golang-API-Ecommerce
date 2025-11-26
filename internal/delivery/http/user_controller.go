package http

import (
	"API-Ecommerce-Evermos/internal/model"
	"API-Ecommerce-Evermos/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	usecase usecase.IUserUsecase
	log     *logrus.Logger
}

func NewUserController(usecase usecase.IUserUsecase, log *logrus.Logger) *UserController {
	return &UserController{usecase, log}
}

func (u UserController) Register(c *fiber.Ctx) error {
	var request model.RegisterUserRequest

	if err := c.BodyParser(&request); err != nil {
		u.log.Warnf("Failed to parse register request: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userResponse, err := u.usecase.Register(request)
	if err != nil {
		u.log.Errorf("Failed to register user: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Register berhasil",
		"data":    userResponse,
	})

}

func (u *UserController) Login(c *fiber.Ctx) error {
	var request model.LoginUserRequest

	if err := c.BodyParser(&request); err != nil {
		u.log.Warnf("Failed to parse login request: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	loginResponse, err := u.usecase.Login(request)
	if err != nil {
		u.log.Warnf("Failed to login: %v", err)
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login berhasil",
		"data":    loginResponse,
	})
}

func (c *UserController) GetCurrent(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	response, err := c.usecase.GetCurrent(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": response,
	})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	var request model.UpdateUserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response, err := c.usecase.Update(userID, request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"data": response,
	})
}
