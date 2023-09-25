package main

import (
	"fiber-apis/database"
	"fiber-apis/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())

	router.SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
