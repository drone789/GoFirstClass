package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func validateAuth(s string) error {
	if s != "123456" {
		return fmt.Errorf("%s", "bad auth token")
	}
	return nil
}

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		h.ServeHTTP(w, r)
	})
}

// authHandler 包装函数，支持链式调用
func authHandler(h http.Handler) http.Handler {
	// 内部利用了适配器函数类型HandlerFunc，将普通的匿名函数转成实现了 http.Handler(接口) 的类型的实例
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := validateAuth(r.Header.Get("auth"))
		if err != nil {
			http.Error(w, "bad auth param", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})

}

func main() {
	http.ListenAndServe(":8080", logHandler(authHandler(http.HandlerFunc(greetings))))
}
