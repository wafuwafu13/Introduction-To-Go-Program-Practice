// リスト6.18
/* 
 1. go get github.com/jinzhu/gorm
 2. go run store.go
*/
package main

import (
  "fmt"
	"github.com/jinzhu/gorm" // 自動マイグレーション機能を持っているのでsetup.sqlが必要ない
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int  // 自動的にこの形式のフィールドは外部キーであると想定して必要な関係を作成する
	CreatedAt time.Time
}

var Db *gorm.DB

// connect to the Db
func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post) // {0 Hello World! Sau Sheong [] 0001-01-01 00:00:00 +0000 UTC}
  
	// Create a post
	Db.Create(&post)
    fmt.Println(post) // {1 Hello World! Sau Sheong [] 2015-04-13 11:38:50.91815604 +0800 SGT}
  
    // Add a comment
	comment := Comment{Content: "いい投稿だね！", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment) // PostIdに手動でアクセスすることはない

    // Get comments from a post
	var readPost Post  
	Db.Where("author = $1", "Sau Sheong").First(&readPost) // 検索結果のレコードを変数readPostに入れる
    var comments []Comment
    Db.Model(&readPost).Related(&comments)
    fmt.Println(comments[0]) // {1 Good post! Joe 1 2015-04-13 11:38:50.920377 +0800 SGT}
}
