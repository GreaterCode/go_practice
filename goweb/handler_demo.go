package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/baidu", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "http://www.baidu.com", http.StatusOK)
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	})

	mux.HandleFunc("/hehe", sayHehe)

	http.ListenAndServe(":8080", mux)
}

func sayHehe(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hehe!!!!!")
}
