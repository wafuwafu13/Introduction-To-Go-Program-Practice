// リスト6.4
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// store data
func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer) // エンコーダを使ってデータをバッファ内にエンコードする
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// load the data
func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw) // 生データからバッファを生成する
	dec := gob.NewDecoder(buffer) // バッファからデコーダを生成する
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead) // {1 Hello World! Sau Sheong}
}
