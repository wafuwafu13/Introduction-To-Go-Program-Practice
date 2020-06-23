// リスト4.4
package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // リクエストを解析
	r.ParseMultipartForm(1024) // 取得するバイト数を指定
	fmt.Fprintln(w, r.Form) // Formフィールドにアクセス  map[hello:[sau sheong world] post:[456] thread:[123]]
	fmt.Fprintln(w, r.MultipartForm) 
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

// 1. ParseFormまたはParseMultipartFormを呼び出して、リクエストを解析する。
// 2. 必要に応じてForm, PostForm, MultipartFormというフィールドから取得する。

// フォームとURLに同一のキーがある場合は常にフォームの値を優先してURLの値より前に配置する

// URLの方のペアを無視したい場合は、r.PostForm を使う

// enctype を multipart/form-data にすると、PostFormフィールドが application/x-www-form-urlencoded しか
// サポートしていないのでURLクエリのキーと値のペアだけを取得する

// application/x-www-form-urlencoded
// fmt.Fprintln(w, r.FormValue("hello"))  sau sheong
// fmt.Fprintln(w, r.PostFormValue("hello"))  sau sheong
// fmt.Fprintln(w, r.PostForm)  map[hello:[sau sheong] post:[456]]
// fmt.Fprintln(w, r.MultipartForm)  <nil>

// multipart/form-data
// fmt.Fprintln(w, r.FormValue("hello"))  world
// fmt.Fprintln(w, r.PostFormValue("hello"))  出力なし
// fmt.Fprintln(w, r.PostForm)  map[]
// fmt.Fprintln(w, r.MultipartForm) &{map[hello:[sau sheong] post:[456]] map[]}  値がファイルだから空になっている