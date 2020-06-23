// リスト4.11
package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str)) // バイト配列を受け取ってHTTPレスポンスのボディに書き込む
}

// curl -i 127.0.0.1:8080/write
// HTTP/1.1 200 OK
// Date: Tue, 23 Jun 2020 00:55:08 GMT
// Content-Length: 95
// Content-Type: text/html; charset=utf-8
// 
// <html>
// <head><title>Go Web Programming</title></head>
// <body><h1>Hello World</h1></body>
// </html>%

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "そのようなサービスはありません。ほかを当たってください")
}

// curl -i 127.0.0.1:8080/writeheader
// HTTP/1.1 501 Not Implemented
// Date: Tue, 23 Jun 2020 00:57:01 GMT
// Content-Length: 82
// Content-Type: text/plain; charset=utf-8
// 
// そのようなサービスはありません。ほかを当たってください

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com") // Location というヘッダとユーザーの転送先にしたい場所の値を追加
	w.WriteHeader(302)
}

// curl -i 127.0.0.1:8080/redirect
// HTTP/1.1 302 Found
// Location: http://google.com
// Date: Tue, 23 Jun 2020 00:59:02 GMT
// Content-Length: 0

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Sau Sheong",
		Threads: []string{"1番目", "2番目", "3番目"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

// curl -i 127.0.0.1:8080/json
// HTTP/1.1 200 OK
// Content-Type: application/json
// Date: Tue, 23 Jun 2020 01:00:43 GMT
// Content-Length: 63
// 
// {"User":"Sau Sheong","Threads":["1番目","2番目","3番目"]}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}

// 渡すのがRequestへのポインタである理由は、ハンドラによるRequestの変更がサーバから見える必要があるので、値渡しではなく参照渡しにする
// ResponseWriterはエクスポートされていない構造体であるresponseへのインターフェースであり、その構造体を値渡しではなく参照渡ししている