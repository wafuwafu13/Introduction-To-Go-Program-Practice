package main

import (
  "testing"
  "time"
)

func TestParallel_1(t *testing.T) {
  t.Parallel() // テストケースを並行実行するため関数Parallelの呼び出し
  time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
  t.Parallel()
  time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
  t.Parallel()
  time.Sleep(3 * time.Second)
}
