package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw  >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

func main() {
	// サイズが3のバッファつきチャネル
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}

// Threw  >> 0
// Threw  >> 1
// Caught << 0
// Caught << 1
// Caught << 2
// Threw  >> 2 // バッファが受け入れを中断するまでに3つの数でいっぱいになっている
// Threw  >> 3
// Threw  >> 4
// Caught << 3
// Caught << 4
