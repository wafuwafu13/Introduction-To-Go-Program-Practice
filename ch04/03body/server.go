// リスト4.3
// http://localhost:8080/body
package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}

// GETリクエストにはメッセージボディがないため、POSTリクエストをサーバに送信する必要がある

// curl -id "first_name=hoge&last_name=wafu" 127.0.0.1:8080/body
// HTTP/1.1 200 OK
// Date: Mon, 22 Jun 2020 23:30:05 GMT
// Content-Length: 31
// Content-Type: text/plain; charset=utf-8

// first_name=hoge&last_name=wafu
