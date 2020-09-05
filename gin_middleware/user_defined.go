package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(MyMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"res": "user define middleware"})
	})

	router.Run(":8080")
}

func MyMiddleware() gin.HandlerFunc {
	//中间件逻辑处理
	fmt.Println("todo....")

	return func(c *gin.Context) {
		fmt.Println("return middleware")
	}
}
