// リスト7.11
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile) // JSONデータからデコーダを生成する
	for { // EOFが検出されるまで繰り返す
		var post Post
		err := decoder.Decode(&post) // JSONデータをデコードし構造体に収納する
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post) // {1 Hello World! {2 Sau Sheong} [{3 Have a great day! Adam} {4 How are you today? Betty}]}
	}
}

// Decoder http.RequestのBodyのように、io.Readerのストリームからデータが入ってくる場合
// Unmarshal 文字列データや、メモリ内に格納されている場合
