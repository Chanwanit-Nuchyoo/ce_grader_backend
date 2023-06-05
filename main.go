package main

import (
	"ce_grader_backend/environment"

	"github.com/gin-gonic/gin"
)

func main() {
	environment.LoadEnvVariables()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
