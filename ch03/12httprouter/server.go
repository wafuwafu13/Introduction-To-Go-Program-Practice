// リスト3.12


package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter" // go get github.com/julienschmidt/httprouter
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) { // Params型には名前付きパラメータが入っていて、メソッドByNameで取得できる
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New() // マルチプレクサを生成
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
