// リスト6.1
package main

import (
	"fmt"
)

// メモリ中に保管するデータ
type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post // ユニークなIDを投稿へのポインタに対応付ける
var PostsByAuthor map[string][]*Post  // 著者名を投稿へのポインタのスライスに対応付ける

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {

	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1]) // &{1 Hello World! Sau Sheong}
	fmt.Println(PostById[2]) // &{2 Bonjour Monde! Pierre}

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)  // &{1 Hello World! Sau Sheong} &{4 Greetings Earthlings! Sau Sheong}
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post) // &{3 Hola Mundo! Pedro}
	}
}
