package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// User represents a user in the system
// User struct to represent a user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Mock database of users
var users = map[string]string{
	"admin": "password123",
	"user":  "admin",
}

// JWT secret key (should be stored securely in a real application)
var jwtSecret = []byte("your-256-bit-secret")

// Register handles user registration
func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Here you would typically add code to save the user to a database
	// For simplicity, we'll just return a success message

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}


// Login handles the login request
func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Check if the user exists and the password is correct
	if storedPassword, exists := users[user.Username]; exists {
		if storedPassword == user.Password {
			// Generate a JWT token
			token, err := generateToken(user.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			// Return the token and a success message
			c.JSON(http.StatusOK, gin.H{
				"message": "User logged in successfully",
				"token":   token,
			})
			return
		}
	}

	// If the user doesn't exist or the password is wrong
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// generateToken creates a JWT token for the given username
func generateToken(username string) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}