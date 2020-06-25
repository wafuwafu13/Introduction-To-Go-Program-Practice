package main

import "testing"
import "time"

func TestPrint1(t *testing.T) {    
  print1()
}

func TestGoPrint1(t *testing.T) {    
  goPrint1()
  time.Sleep(1 * time.Millisecond) // 終了する前に出力を生成できる
}
