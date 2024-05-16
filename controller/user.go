package controller

import (
	"log"
	"time"

	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
	"github.com/Sahil-Sayyad/Trademarkia/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Create User - Registers a new user
func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	user.Password = string(hashedPassword)

	//Create user in the database
	result := database.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(user)
}

// Login - Authenticate a user and returns a JWT token
func Login(c *fiber.Ctx) error {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//Find user by email
	var user model.User

	if result := database.DB.Where("email=?", input.Email).First(&user); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	//Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})

	}

	//Generate JWT token
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})

	}

	cookie := fiber.Cookie{
		Name:     "jwt-user",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), //token expires in 24 hours
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"token": token})

}

// Create Admin - Registers a new admin
func CreateAdmin(c *fiber.Ctx) error {
	var admin model.Admin

	if err := c.BodyParser(&admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	admin.Password = string(hashedPassword)

	//Create user in the database
	result := database.DB.Create(&admin)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(admin)
}

// Login - Authenticate a admin and returns a JWT token
func LoginAdmin(c *fiber.Ctx) error {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//Find user by email
	var admin model.Admin

	if result := database.DB.Where("email=?", input.Email).First(&admin); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	//Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})

	}

	//Generate JWT token
	token, err := utils.GenerateTokenAdmin(admin.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})

	}

	cookie := fiber.Cookie{
		Name:     "jwt-admin",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour), //token expires in 24 hours
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"token": token})

}
// View user's order history
func GetUserDashboard(c *fiber.Ctx) error {

    // Get the authenticated user's ID from the context
    userID, ok := c.Locals("userID").(string)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Unauthorized",
        })
    }
	var orders []model.Order
    result := database.DB.Preload("Products").Where("user_id = ?", userID).Find(&orders)

    if result.Error != nil {
        // Log the error for debugging
        log.Println("Error fetching orders:", result.Error) 
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch orders",
        })
    }

    return c.JSON(orders) 
}
