package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
	"path"
	"strconv"
)

func main() {
	// connect to the Db
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: db})) // 構造体Postを渡してhandleRequestを登録
	server.ListenAndServe()
}


// main handler function
func handleRequest(t Text) http.HandlerFunc { // インターフェイスTextを渡す
	return func(w http.ResponseWriter, r *http.Request) { // 正しいシグネチャの関数を返す
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t) // 実際のハンドラにインタフェースTextを渡す
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Retrieve a post
// GET /post/1
func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) { // インタフェースTextを受け入れ
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id) // 構造体Postからデータを取得
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Create a post
// POST /post/
func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// Update a post
// PUT /post/1
func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// Delete a post
// DELETE /post/1
func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
