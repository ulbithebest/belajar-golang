package belajarHandler

import (
	"belajar-golang/database"
	"belajar-golang/internal/model"

	"github.com/gofiber/fiber/v2"
)

// GetUsers godoc
// @Summary Get all users
// @Description Menampilkan semua data user
// @Tags Users
// @Accept application/json
// @Produce json
// @Success 200 {object} model.User
// @Router /api/belajar [get]
func GetUsers(c *fiber.Ctx) error {
	var users []model.User

	// Find all users in database
	result := database.DB.Find(&users)

	// Check for errors during query execution
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return list of users
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data User Berhasil Ditampilkan!",
		"data":    users,
	})
}

// CreateUser godoc
// @Summary insert data user.
// @Description get data user.
// @Tags Users
// @Accept application/json
// @Param request body model.User true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.User
// @Router /api/belajar [post]
func CreateUser(c *fiber.Ctx) error {
	// Parse request body
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// Insert new user into database
	result := database.DB.Create(&user)

	// Check for errors during insertion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return success message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Berhasil Ditambahkan!",
		"data":    user,
	})
}

// GetUser godoc
// @Summary Data user berdasarkan ID.
// @Description get data user.
// @Tags Users
// @Accept application/json
// @Param id_user path string true "Masukan ID User"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/belajar/{id_user} [get]
func GetUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")

	// Find user by id_user in database
	var user model.User
	result := database.DB.First(&user, id)

	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditermukan!",
		})
	}

	// Check for errors during query
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return user
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    user,
	})
}

// UpdateUser godoc
// @Summary update data user.
// @Description update data user.
// @Tags Users
// @Accept application/json
// @Param id_user path string true "Masukan ID User"
// @Param request body model.User true "Payload body [RAW]"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/belajar/{id_user} [put]
func UpdateUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")

	// Find user by id_user in database
	var user model.User
	result := database.DB.First(&user, id)

	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan",
		})
	}

	// Parse request body
	var newUser model.User
	if err := c.BodyParser(&newUser); err != nil {
		return err
	}

	// Update user in database
	result = database.DB.Model(&user).Updates(newUser)

	// Check for errors during update
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Berhasil Diperbarui!",
		"data":    user,
	})
}

// DeleteUser godoc
// @Summary Hapus Data user berdasarkan ID.
// @Description delete data user.
// @Tags Users
// @Accept application/json
// @Param id_user path string true "Masukan ID User"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/belajar/{id_user} [delete]
func DeleteUser(c *fiber.Ctx) error {
	// Get id_user parameter from request url
	id := c.Params("id_user")

	// Find user by id_user in database
	var user model.User
	result := database.DB.First(&user, id)

	// Check if user exists
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User Tidak Ditemukan",
		})
	}

	// Delete user from database
	result = database.DB.Delete(&user)

	// Check for errors during deletion
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User Berhasil Dihapus!",
		"data":    user,
	})
}
