package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/maspratama/go-patungan/config"
)

func main() {
	// connect to db
	db.ConncetToDB()

	app := fiber.New()

	app.Listen(":8082")
}
