package main

import (
	"fmt"
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
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		}
	}
}

// Hola Mundo! from B
// Hello World! from A
// fatal error: all goroutines are asleep - deadlock!

// ゴルーチンがあるチャネルから値を取り出すと、そのチャネルから値を取り出す他のゴルーチンが全てブロックされスリープしてしまう