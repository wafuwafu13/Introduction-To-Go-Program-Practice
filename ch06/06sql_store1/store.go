// リスト6.6
/*
 1. Postgresを起動
 2. createuser -P -d gwp
    （passwd gwp）
 3. creatdb gwp
 4. psql -U gwp -f setup.sql -d gwp
   （一番最初は ERRORが表示されるが、そのままでOK)
 5. go get "github.com/lib/pq"
 6. go run store.go
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQLのドライバ initで自身を登録する
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

// connect to the Db 全てのパッケージに対して自動的に呼び出される
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// get all posts
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() { // sql.Rowを返し、行がなくなるとio.EOFを返す
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author) // 返された結果の値をからの構造体Postにコピーできる
	return
}

// Create a new post
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id" // プリペアドステートメントの定義 idを返すように指示
	stmt, err := Db.Prepare(statement) // sql.Stmtへの参照を生成
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id) // プリペアドステートメントを実行 レシーバであるpostのIdフィールドに対して、SQLクエリにより返されるidフィールドの値が設定される
	return
}

// Update a post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author) // グローバルDBのメソッドExecを使う
	return
}

// Delete a post
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

// Delete all posts
func DeleteAll() (err error) {
	_, err = Db.Exec("delete from posts")
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"} // Idフィールドはデータベースによって自動的に設定される

	// Create a post
	fmt.Println(post) // {0 Hello World! Sau Sheong}
	post.Create()
	fmt.Println(post) // {1 Hello World! Sau Sheong}

	// Get one post
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost) // {1 Hello World! Sau Sheong}

	// Update the post
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	// Get all posts
	posts, _ := Posts(10)
	fmt.Println(posts) // [{1 Bonjour Monde! Pierre}]

	// Delete the post
	readPost.Delete()

	// Get all posts
	posts, _ = Posts(10)
	fmt.Println(posts) // []

	// Delete all posts
  // DeleteAll()
}
