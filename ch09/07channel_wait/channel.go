// リスト9.7
package main

import "fmt"
import "time"

func printNumbers2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	w <- true // チャネルにブール値を入れて中断を解除する
}

func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers2(w1)
	go printLetters2(w2)
	<-w1 // 何かが入るまでチャネルは実行を中断する　ゴルーチンが終了したときにプログラムの中断を解除することにしか関心がない
	<-w2
	fmt.Printf("\n")
}
