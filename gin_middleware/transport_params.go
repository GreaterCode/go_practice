package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Middleware1(c *gin.Context) {
	c.Set("key1", 123)
	c.Set("key2", "renwoxing")

}
func main() {
	router := gin.New()
	router.GET("/", Middleware1, func(c *gin.Context) {
		key1 := c.GetInt("key1")
		key2 := c.GetString("key2")
		c.JSON(http.StatusOK, gin.H{
			"key1": key1,
			"key2": key2,
		})
	})
	router.Run(":8080")
}
