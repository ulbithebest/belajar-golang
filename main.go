package main

import (
	"belajar-golang/database"
	_ "belajar-golang/docs"
	"belajar-golang/router"

	"github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Belajar Golang API
// @description This is a sample API for Belajar Golang
// @version 1
// @host localhost:3000
// @BasePath /
func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	// Setup the router
	router.SetupRoutes(app)

	// Generate Swagger documentation
	app.Get("/docs/*", swagger.HandlerDefault)

	// Listen on PORT 3000
	app.Listen(":3000")
}
