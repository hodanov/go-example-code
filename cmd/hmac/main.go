package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// サーバーサイドのDBがあるとして、その中に登録されているユーザの情報と仮定
var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

// サーバーサイドでユーザの判定
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

func main() {
	// サーバーが提供しているapiKeyをクライアントが持っていると仮定
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("data")

	// sha256でハッシュを作る
	h := hmac.New(sha256.New, []byte(apiSecret))

	// 何か値を付け加える
	h.Write([]byte("data"))

	// エンコードして表示できるようにする
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign)

	// apiKeyとapiSecretが正しいかどうか、サーバーサイドに問い合わせる
	Server(apiKey, sign, data)

}
