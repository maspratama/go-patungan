package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	db "github.com/maspratama/go-patungan/config"
	"github.com/maspratama/go-patungan/routes"
)

func main() {
	// connect to db
	db.ConncetToDB()

	app := fiber.New()
	routes.Setup(app)

	err := app.Listen(":8082")
	if err != nil {
		fmt.Printf("Server port error")
	}
}
