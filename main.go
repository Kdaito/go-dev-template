package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongpongpong",
		})
	})

	if err := r.Run(); err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
