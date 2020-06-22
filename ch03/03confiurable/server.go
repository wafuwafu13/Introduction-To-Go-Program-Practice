package main

import (
	"net/http"
)

func main() {
	server := http.Server{ // Serverという構造体があり、サーバーの設定を変更できる
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
