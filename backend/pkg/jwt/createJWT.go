package jwt

import (
	"encoding/base64"
	"log"
)

func generateJSON(info *UserInfo) {
}

func encodeJWT(s string) {
	_, err := base64.URLEncoding.DecodeString(s)

	if err != nil {
		log.Println("Base64 error:", err)
	}
	return
}
