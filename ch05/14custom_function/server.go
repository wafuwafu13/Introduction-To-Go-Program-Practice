// リスト5.14
package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate} // fdate という名前を関数formatDateに対応づける
	t := template.New("tmpl.html").Funcs(funcMap) // fumcMapがテンプレートに付加され、テンプレートファイルtmpl.htmlの解析で使用される
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
