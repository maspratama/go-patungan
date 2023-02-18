package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	db "github.com/maspratama/go-patungan/config"
	"github.com/maspratama/go-patungan/user"
)

func main() {
	// connect to db
	db.ConncetToDB()

	app := fiber.New()

	userRepository := user.NewRepository(db.DB)
	user := user.Users{
		Name:           "Mastama",
		Email:          "mastama@gmail.com",
		Occupation:     "Middleware Developer",
		PasswordHash:   "mastama456",
		AvatarFileName: "image.png",
		Role:           "Admin",
	}
	_, err := userRepository.Save(user)
	if err != nil {
		return
	}

	err = app.Listen(":8082")
	if err != nil {
		fmt.Printf("Server port error")
	}
}
