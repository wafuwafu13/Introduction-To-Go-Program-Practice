// リスト7.2
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`  // XML要素の名前自体を保存
	Id      string   `xml:"id,attr"` // 属性を保存
	Content string   `xml:"content"` // モードフラグがない場合、構造体のフィールドは構造体名と同じ名前のXML要素と対応づけられる
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"` // 未処理のXMLを保存
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"` // 文字データを保存
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
	fmt.Println(post)
}

// {{ post} 1 Hello World! {2 Sau Sheong}  構造体 id content id name
//   <content>Hello World!</content>  　　内部のXMLを単に出力したもの
//   <author id="2">Sau Sheong</author>
// }

