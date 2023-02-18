package routes

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/maspratama/go-patungan/config"
	"github.com/maspratama/go-patungan/controller"
	"github.com/maspratama/go-patungan/user"
)

func Setup(app *fiber.App) {
	//user routes

	userRepository := user.NewRepository(db.DB)

	userService := user.NewService(userRepository)

	userController := controller.NewUserController(userService)

	api := app.Group("/api/v1")

	api.Post("/users", userController.CreateUserRegister)
}
