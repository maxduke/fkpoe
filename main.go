package main

import (
	"fkpoe/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		if os.IsNotExist(err) {
			log.Info(".env file does not exist")
		} else {
			log.Fatal("Error loading .env file")
		}
	}
	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Forwarded-For",
	})
	router.SetupRoutes(app)
	app.Listen(":3651")
}
