// リスト7.5など
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"` // XML要素commentsを飛ばして下位要素commentを直接取得
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post.XMLName.Local) // post
	fmt.Println(post.Id) // 1
	fmt.Println(post.Content) // Hello World!
	fmt.Println(post.Author) // {2 Sau Sheong}
	fmt.Println(post.Xml) // そのまま出力される
	fmt.Println(post.Author.Id) // 2
	fmt.Println(post.Author.Name) // Sau Sheong
	fmt.Println(post.Comments) // [{1 Have a great day! { Adam}} {2 How are you today? { Betty}}]
	fmt.Println(post.Comments[0].Id) // 1
	fmt.Println(post.Comments[0].Content) // Have a great day!
	fmt.Println(post.Comments[0].Author) // { Adam}
	fmt.Println(post.Comments[1].Id)  // 2
	fmt.Println(post.Comments[1].Content) // How are you today?
	fmt.Println(post.Comments[1].Author) // { Betty}
}
