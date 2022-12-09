package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	log.Fatal(err)
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("hello world"))
}
