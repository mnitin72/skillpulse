package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	r.GET("/api/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"skills": 5,
			"hours": 120,
		})
	})

	r.Run(":8080")
}
