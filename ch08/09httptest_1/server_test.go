package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux() // テストを実行するマルチプレクサを生成
	mux.HandleFunc("/post/", handleRequest) // テスト対象のハンドラを付加

	writer := httptest.NewRecorder() // 返されたHTTPレスポンスを取得
	request, _ := http.NewRequest("GET", "/post/1", nil) // テストしたいハンドラ宛のリクエストを作成
	mux.ServeHTTP(writer, request) // テスト対象のハンドラにリクエストを送信

	if writer.Code != 200 { // ResponseRecoderにより結果をチェックs
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
