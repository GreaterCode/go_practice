package main

import (
	"log"
	"net/http"
)

/**
    FileServer：
    1.www.xx.com/ 根路径 直接使用

　　　　　http.Handle("/", http.FileServer(http.Dir("/tmp")))

　　2.www.xx.com/c/ 带有请求路径的 需要添加函数
　　　　　　http.Handle("/c/", http.StripPrefix("/c/", http.FileServer(http.Dir("/tmp"))))
*/
func main() {
	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./")))
	if err != nil {
		log.Fatal(err)
	}

}
