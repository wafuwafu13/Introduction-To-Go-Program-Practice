// リスト4.16
package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1) // first_cookie="Go Web Programming"
	fmt.Fprintln(w, cs) // [first_cookie="Go Web Programming" second_cookie="Manning Publications Co"]
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}

// クッキーを設定するときにExpiresやMaxAgeのフィールドを設定しなかったので、返されるのはセッションクッキー。
// ブラウザを終了して get_cookie したら何も返されない