package main_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "../25test_ginkgo"
//	. "gwp/Chapter_8_Testing_Web_Applications/test_ginkgo"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Get a post", func() { // 関数initの呼び出しを避けるための _  ユーザーストーリー
	var mux *http.ServeMux
	var post *FakePost
	var writer *httptest.ResponseRecorder

	BeforeEach(func() {
		post = &FakePost{}
		mux = http.NewServeMux()
		mux.HandleFunc("/post/", HandleRequest(post)) // mainからエクスポートされ、ここでテストされる
		writer = httptest.NewRecorder()
	})

	Context("using an id", func() { // シナリオ1
		It("should get a post", func() {
			request, _ := http.NewRequest("GET", "/post/1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(200)) // Gomegaマッチャーを使用

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)

			Expect(post.Id).To(Equal(1))
		})
	})

	Context("using a non-integer id", func() { // シナリオ2
		It("should get a HTTP 500 response", func() {
			request, _ := http.NewRequest("GET", "/post/hello", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(500))
		})
	})

})
