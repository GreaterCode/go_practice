package main

import "net/http"
import _ "net/http/pprof"

func main() {
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()
}
