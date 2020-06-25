package main

import "fmt"
import "time"
import "sync"

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done() // カウンタを減算
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done() // カウンタを減算
}

func main() {
	var wg sync.WaitGroup // WaitGroupの宣言
	wg.Add(2) // カウンタの初期化
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait() // カウンタが0になるまで実行を中断
	fmt.Printf("\n")
}
