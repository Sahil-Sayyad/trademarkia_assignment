package controller

import (
	"time"

	"github.com/Sahil-Sayyad/Trademarkia/database"
	"github.com/Sahil-Sayyad/Trademarkia/model"
	"github.com/Sahil-Sayyad/Trademarkia/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Create User - Registers a new user
func CreateUser(c *fiber.Ctx) error{
	var user model.User

	if err := c.BodyParser(&user);  err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON"})
	}

	//Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Could not hash password"})
	}
	user.Password = string(hashedPassword)

	//Create user in the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(user)
}

//Login - Authenticate a user and returns a JWT token
func Login(c *fiber.Ctx) error {
	
	var input struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON"})
	} 

	//Find user by email
	var user model.User

	if result := database.DB.Where("email=?",input.Email).First(&user); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid credentials"})
	}

	//Compare password 
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid credentials"})

	}
   
	//Generate JWT token
	token , err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})

	}

	cookie := fiber.Cookie{
		Name : "jwt",
		Value: token,
		Expires: time.Now().Add(24 * time.Hour), //token expires in 24 hours
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{"token": token})

}