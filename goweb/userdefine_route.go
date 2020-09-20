package main

import (
	"fmt"
	"net/http"
)

// User define ropute
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloRoute(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello! my route！！！")

}

func main() {
	myMux := &MyMux{}
	http.ListenAndServe(":8080", myMux)
}
