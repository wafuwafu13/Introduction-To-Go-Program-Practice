package main

import (
	"crypto/rand"
	"crypto/rsa" // 秘密鍵生成
	"crypto/x509" // 証明書生成
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// SSL証明書は公開鍵などの情報を含み、X.509形式にまとめられた一片のデータでWebサーバに保管される
// クライアントがサーバにリクエストを出すと、サーバから証明書が返され、
// クライアントが本物であるか確かめたあとでランダムな鍵を生成し、証明書(証明書内の公開鍵)を使ってその鍵を暗号化する

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 証明書の構成を設定するための構造体
	template := x509.Certificate{
		SerialNumber: serialNumber, // 認証局によって発行される一意の番号(ランダムなとても大きい整数)
		Subject:      subject, // 識別名
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 有効期間
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // サーバ認証に使用することを明示
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, // サーバ認証に使用することを明示
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")}, // 効力を限定
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048) // 秘密鍵を生成

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) // バイトデータのスライスを生成
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}) // 証明書データを符号化
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)}) // 秘密鍵を符号化
	keyOut.Close()
}
