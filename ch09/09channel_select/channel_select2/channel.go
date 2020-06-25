package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello World!"
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond) // select文を呼び出すのが早すぎて、チャネルからきちんと値を取り出せないうちにループが終了してしまうのを防ぐ
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Println("Default")
		}
	}
}

// Hello World! from A
// Hola Mundo! from B
// Default
// Default
// Default
