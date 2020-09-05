package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func statCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("test", "123456")
		// 请求前逻辑
		c.Next()
		//请求后逻辑
		latency := time.Since(t)
		log.Printf("total spend time: %d us", latency/1000)

	}

}

func main() {
	router := gin.Default()
	router.Use(statCost())
	router.GET("/", func(c *gin.Context) {
		test := c.MustGet("test").(string)
		log.Println(test)
		c.JSON(http.StatusOK, gin.H{"msh": "success"})
	})
	router.Run(":8080")
}
