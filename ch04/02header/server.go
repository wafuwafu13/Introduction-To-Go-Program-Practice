// リスト4.2
/* 
 1. go run server.go
 2. http://localhost:8080/headers  を表示
*/

package main

import (
	"fmt"
	"net/http"
)

// ヘッダは「キーが文字列型」で「値が文字列型のスライス」のマップ

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
 http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}

// h := r.Header["Accept-Encoding"] // [gzip, deflate]
