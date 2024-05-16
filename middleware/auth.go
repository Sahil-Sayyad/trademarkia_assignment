package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("sahil")

type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func UserauthMiddleware(c *fiber.Ctx) error {
	var tokenString string

	cookie := c.Cookies("jwt-user")

	if cookie != "" {
		tokenString = cookie
	}
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "login or sign-up  "})
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token signature"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		c.Locals("userID", claims.UserID)
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
}

func AdminauthMiddleware(c *fiber.Ctx) error {
	var tokenString string

	cookie := c.Cookies("jwt-admin")

	if cookie != "" {
		tokenString = cookie
	}
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Admin access only "})
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token signature"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		c.Locals("userID", claims.UserID)
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
}
