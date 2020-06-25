package main

import "testing"
import "time"

func TestPrint1(t *testing.T) {    
  print1()
}

func TestGoPrint1(t *testing.T) {    
  goPrint1()
  time.Sleep(1 * time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
	goPrint2() // printNumbers2 と printLetters2 が独立して実行されていて、出力を争う
	time.Sleep(1 * time.Millisecond)
}
