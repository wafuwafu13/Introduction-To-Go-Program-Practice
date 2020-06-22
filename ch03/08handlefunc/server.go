package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handlerは指定しない -> DefaultServeMuxをハンドラとして利用
	}
	http.HandleFunc("/hello", hello) // 関数helloをハンドラに変換して、DefaultServeMuxに登録
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}

// func HandleFunc(pattern string, handler func(ResponseWriter *Request)) {
// 	DefaultServeMux.HandleFunc(pattern, handler)
// }
// 
// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter *Request)) {
// 	mux.Handle(pattern, HandleFunc(handler))
// }