package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	//冒号:加上一个参数名组成路由参数
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, name)

	})

	//*号能匹配的规则就更多
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// URL参数：URL 参数通过 DefaultQuery 或 Query 方法获取。
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "guest") //设置默认值
		c.String(http.StatusOK, fmt.Sprintf("Hello %s ", name))
	})

	// 表单参数，常见格式有四种：
	//application/json，application/x-www-form-urlencoded, application/xml和multipart/form-data
	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert")
		username := c.PostForm("username")
		password := c.PostForm("password")
		hobbies := c.PostFormArray("hobby")
		c.String(http.StatusOK, fmt.Sprintf("type:%s, username:%s, password:%s, hobbys:%v",
			type1, username, password, hobbies))

	})

	// 上传文件
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, file.Filename)

		/*
			也可以直接用io操作，拷贝文件数据
			        out, err := os.Create(filename)
			        defer out.Close()
			        _, err = io.Copy(out, file)
		*/
		c.String(http.StatusOK, fmt.Sprintf("%s is upload", file.Filename))
	})

	// 上传多个文件
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload_multi", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err:%s", err.Error()))
			return

		}
		files := form.File["files"]
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}

		}
		c.String(http.StatusOK, fmt.Sprintf("upload successfully %d files", len(files)))

	})
	// 自定义HTTP服务器配置
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
