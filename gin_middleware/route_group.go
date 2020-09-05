package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1", gin.Logger(), gin.Recovery())
	{
		v1.GET("/test1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"res": "test1"})
		})
		v1.GET("/test2", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"res": "test2"})
		})
	}

	router.Run(":8080")
}
