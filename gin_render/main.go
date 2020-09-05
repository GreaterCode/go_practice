package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}

		msg.Name = "renwoxing"
		msg.Message = "hello"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})

	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hello", "status": http.StatusOK})
	})

	router.GET("/protobuf", func(c *gin.Context) {
		format := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  format,
		}
		c.ProtoBuf(http.StatusOK, data)

	})

	// html模板
	//加载模板
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "main website"})
	})

	// 加载静态文件
	router.StaticFS("/Users/yuanting/books", http.Dir("."))
	router.StaticFS("/usr", http.Dir("/bin"))
	router.StaticFile("/bin", ".")

	// 重定向
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")

	})
	router.Run(":8080")

}
