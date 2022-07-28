package main

import (
	"fmt"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greetings)) // 将greetings 转成 HandlerFunc()类型
}

// $GOROOT/src/net/http/server.go
//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
//
//type HandlerFunc func(ResponseWriter, *Request)
//
//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//	f(w, r)
//}
