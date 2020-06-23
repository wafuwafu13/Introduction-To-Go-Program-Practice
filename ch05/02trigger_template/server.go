// リスト5.2
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html") // テンプレートファイル tmpl.html を解析
	t.Execute(w, "Hello World!") // データをテンプレートに当てはめる
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
