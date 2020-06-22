// リスト3.11
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

// Hendlerを受け取ってハンドラを返す
func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h) // Handler called - main.HelloHandler　%vの場合 - {}
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// some code to make sure the user is authorized　ハンドラが呼び出される前に実行される
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
}
