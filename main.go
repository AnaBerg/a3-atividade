package main

import (
	"fmt"
	"go-atividade-a3/address"
	"go-atividade-a3/client"

	"go-atividade-a3/database"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	client.Routes(app)
	address.Routes(app)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.DBConn.AutoMigrate(client.Client{}, address.Address{})

	fmt.Println("Database connection successfully opened")

}

func main() {
	app := fiber.New()
	setupRoutes(app)

	initDatabase()

	app.Listen(8080)
}
