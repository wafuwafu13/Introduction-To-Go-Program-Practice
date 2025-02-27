// リスト6.2
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")

	// ファイルへの書き込みとファイルからの読み込み。WriteFileとReadFileを利用
	err := ioutil.WriteFile("data1", data, 0644) // ファイル名, 書き込むデータ, ファイルに設定するパーミッション
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1)) // Hello World!

	// File構造体を利用したファイルの読み書き
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes) // Wrote 13 bytes to file

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes) // Read 13 bytes from file
	fmt.Println(string(read2)) // Hello World!
}
