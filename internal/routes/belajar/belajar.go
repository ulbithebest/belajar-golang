package belajarRoutes

import (
	belajarHandler "belajar-golang/internal/handlers/belajar"

	"github.com/gofiber/fiber/v2"
)

func SetupBelajarRoutes(router fiber.Router) {
	user := router.Group("/belajar")
	// Create a user
	user.Post("/", belajarHandler.CreateUser)
	// Read all users
	user.Get("/", belajarHandler.GetUsers)
	// // Read one user
	user.Get("/:id_user", belajarHandler.GetUser)
	// // Update one user
	user.Put("/:id_user", belajarHandler.UpdateUser)
	// // Delete one user
	user.Delete("/:id_user", belajarHandler.DeleteUser)
}
