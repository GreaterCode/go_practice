package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	fmt.Println("Form", r.Form)
	fmt.Println("method: ", r.Method)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("UserAgent: ", r.UserAgent())

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("v: ", v)
	}
	fmt.Fprintf(w, "hello world")

}

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)

	}
}
