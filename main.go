package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/myapi", func(c *gin.Context) {
		data := c.Query("data")
		c.JSON(200, gin.H{
			"data": data,
		})
	})
	r.Run(":8081")
}