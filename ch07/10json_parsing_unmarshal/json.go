// リスト7.10
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int       `json:"id"` // JSONデータのキーidの値を構造体Postのフィールドidに対応付ける
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
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	fmt.Println(string(jsonData))
	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post.Id) // 1
	fmt.Println(post.Content) // Hello World!
	fmt.Println(post.Author.Id) // 2
	fmt.Println(post.Author.Name) // Sau Sheong
	fmt.Println(post.Comments[0].Id) // 1
	fmt.Println(post.Comments[0].Content) // Have a great day!
	fmt.Println(post.Comments[0].Author) // Adam

}
