package main

import (
	"fmt"
	"net/http"
)

// ハンドラ = ServerHttp(http.ResponseWriter, *http.Request)というシグネチャのメソッドを持つもの

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler, // ハンドラをサーバに割り当てる　マルチプレクサを使っておらず、サーバに届く全てのリクエストがこのハンドラに行く(Hello Worldを返す以外何もしない)
	}
	server.ListenAndServe()
}
