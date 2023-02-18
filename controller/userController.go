package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maspratama/go-patungan/helper"
	user "github.com/maspratama/go-patungan/user"
)

type userController struct {
	userService user.Service
}

func NewUserController(userService user.Service) *userController {
	return &userController{userService}
}

func (h *userController) CreateUserRegister(ctx *fiber.Ctx) error {
	var input user.RegisterUserInput

	err := ctx.BodyParser(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := fiber.Map{"errors": errors}

		response := helper.APIResponse("register account failed", 422, "errors", errorMessage)
		ctx.JSON(response)
		return nil
	}

	newUser, err := h.userService.CreateUserRegister(input)

	if err != nil {
		response := helper.APIResponse("register account failed", 400, "error", nil)
		ctx.JSON(response)
		return nil
	}

	formatter := user.FormatUser(newUser, "tokentokentoken")
	response := helper.APIResponse("account has been registerd", 200, "success", formatter)
	ctx.JSON(response)

	return nil
}
