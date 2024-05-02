package main

import (
	"Catsocial/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	setupRouter(r)
	r.Run(":8080")
}

func setupRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		users := v1.Group("/user")
		{
			users.POST("/register", handler.RegisterUser)
			users.POST("/login", handler.LoginUser)
		}
		cat := v1.Group("/cat")
		{
			cat.POST("/", handler.AddCat)
			cat.GET("/", handler.GetCat)
		}
	}
}
