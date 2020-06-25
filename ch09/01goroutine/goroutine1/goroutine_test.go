package main

import "testing"

func TestPrint1(t *testing.T) {    
  print1()
}

func TestGoPrint1(t *testing.T) {    
  goPrint1() // ゴルーチンが何か出力できる前にテストケースが終了する
}
