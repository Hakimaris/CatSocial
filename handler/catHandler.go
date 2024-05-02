package handler

import (
	"github.com/gin-gonic/gin"
)

// Handler register
func AddCat(c *gin.Context) {
	// Logic for registering a user
	c.JSON(200, gin.H{"message": "kat + 1"})
}

// Handler login
func GetCat(c *gin.Context) {
	// Logic for user login
	c.JSON(200, gin.H{"message": "KITR :D"})
}
