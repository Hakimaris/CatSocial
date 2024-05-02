package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "time"
)
// RegisterRequest represents the JSON request body for user registration
type RegisterRequest struct {
	Email    string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`

    // Add other fields as needed
}

// Handler register
func RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request fields (you can add more validation logic here if needed)

	// Assuming validation is successful, generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// claims["email"] = req.Email
	// claims["name"] = req.Name
	// claims["password"] = req.Password


	// Set token expiration (e.g., 8 hours)
	claims["exp"] = time.Now().Add(8 * time.Hour).Unix()

	// Sign the token with a secret
	// Note: Replace "your-secret" with your actual secret key
	tokenString, err := token.SignedString([]byte("your-secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Prepare response JSON
	response := gin.H{
		"message": "User registered successfully",
		"data": gin.H{
			"email":       req.Email,
			"name":        req.Name,
			"accessToken": tokenString,
		},
	}

	c.JSON(200,response)
}

// Handler login
func LoginUser(c *gin.Context) {
	// Logic for user login
	c.JSON(200, gin.H{"message": "p log"})
}
