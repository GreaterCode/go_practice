package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", gin.Recovery(), gin.Logger(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "single route"})
	})

	router.Run(":8080")
}
