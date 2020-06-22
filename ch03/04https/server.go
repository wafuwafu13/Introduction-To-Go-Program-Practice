/* https://localhost:8080/ にアクセスしてください */

package main

import (
	"net/http"
)

// httpsは、SSLの上にHTTPのレイヤーを置いただけのものであり、クライアントとサーバの間の通信を暗号化している

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem") // SSL証明書(認証局or自作)、サーバ用の秘密鍵
}
