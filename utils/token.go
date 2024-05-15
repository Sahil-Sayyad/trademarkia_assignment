package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken generates a JWT token with the given userID
func GenerateToken(userID uint) (string, error) {

    // Get secret key from environment variable
    secretKey := os.Getenv("JWT_SECRET_KEY")

    if secretKey == "" {
        return "", errors.New("JWT_SECRET_KEY environment variable not set")
    }

    // Set claims 
    claims := jwt.MapClaims{
        "userId": strconv.Itoa(int(userID)), // Convert userID to string
        "exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    // Create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString([]byte(secretKey))
    return tokenString, err
}
