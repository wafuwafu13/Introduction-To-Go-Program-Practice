// リスト7.7
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}
  // output, err := xml.Marshal(&post)
  output, err := xml.MarshalIndent(&post, "", "\t\t") // 構造体を組み替えて(marshal)バイト列のXMLデータにする 各行の先頭につけるプレフィックス, インデント文字
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = ioutil.WriteFile("post.xml", []byte(xml.Header + string(output)), 0644) // XML宣言
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}

}
