package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟数据
var secrets = gin.H{
	"zhangsan": gin.H{"email": "zhangsan@example.com", "phone": "123456"},
	"lisi":     gin.H{"email": "lisi@example.com", "phone": "123456"},
	"wangwu":   gin.H{"email": "wangwu@example.com", "phone": "123456"},
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, secrets)
	})
	// 为/admin路由组设置auth, 路由组使用 gin.BasicAuth() 中间件
	auth := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"zhangsan": "zhangsan",
		"lisi":     "lisi",
		"wangwu":   "wangwu",
	}))

	// /admin/secrets 端点
	auth.GET("/secrets", func(c *gin.Context) {
		//获取用户名
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "No secret"})
		}

	})
	router.Run(":8080")
}
