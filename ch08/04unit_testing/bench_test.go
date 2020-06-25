package main

import (
  "testing"
)

// ベンチマークを取りたいコードをb.N回繰り返し、実行時間のベンチマークの精度をあげる
func BenchmarkDecode(b *testing.B) {
  for i := 0; i < b.N; i++ {
    decode("post.json") 
  }
}

func BenchmarkUnmarshal(b *testing.B) {
  for i := 0; i < b.N; i++ {
    unmarshal("post.json")
  }
}