package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shakezidin/model"
)

func main() {
	r := gin.Default()

	requestChannel := make(chan model.Request)

	resultChannel := make(chan model.ConvertedMessage)

	r.POST("/submit", func(c *gin.Context) {
		var req model.Request

		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON format"})
			return
		}

		requestChannel <- req

		converted := <-resultChannel

		c.JSON(200, gin.H{
			"Status":    "Success",
			"message":   "Request received and processed by the worker",
			"converted": converted})
	})

	r.Run(":8080")
}

