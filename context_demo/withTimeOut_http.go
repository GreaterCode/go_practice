package main

import (
	"context"
	"net/http"
	"time"
)

type IndexHandler struct { }

func (this *IndexHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	if r.URL.Query().Get("count") == "" {
		w.Write([]byte("this is header"))
	}else {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()
		c := make(chan string)
		go CountData(c)
		select {
			case <- ctx.Done():
				w.Write([]byte("timeout"))
			case ret := <- c:
				w.Write([]byte(ret))
		}
	}
}

func CountData(c chan string) chan string {
	time.Sleep(time.Second * 5)
	c <- "states result:"
	return c
}



func main() {
	mux := http.NewServeMux()
	mux.Handle("/", new(IndexHandler))
	http.ListenAndServe(":8082", mux)
}
