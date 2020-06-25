package main

import "testing"
import "time"

// CPU増やしても変化なし
func BenchmarkPrint1(b *testing.B) {    
  for i := 0; i < b.N; i++ {
    print1()
  }
}

// CPU2個で速くなるがCPU4個で悪化 高いコスト見合った並行化の効果が得られない場合、かえって実行速度は低下する
func BenchmarkGoPrint1(b *testing.B) {    
  for i := 0; i < b.N; i++ {
    goPrint1()
  }
}

// CPU増やしても変化なし
func BenchmarkPrint2(b *testing.B) {   
  for i := 0; i < b.N; i++ {
    printLetters2()
  }
}

// CPU2,4個で速くなる
func BenchmarkGoPrint2(b *testing.B) {   
  for i := 0; i < b.N; i++ {
    goPrint2()
  }
}


func TestPrint1(t *testing.T) {    
  print1()
}

func TestGoPrint1(t *testing.T) {    
  goPrint1()
  time.Sleep(1 * time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
    goPrint2()
    time.Sleep(1 * time.Millisecond)
}
