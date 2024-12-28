package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func generateJSON(info *UserInfo) (string, error) {
	headers := &JWTHeaders{
		Alg: "HS256", // HMAC SHA-256
		Typ: "JWT",
	}

	date := time.Now().AddDate(0, 0, 1).Unix()
	claims := &JWTBody{
		Iss: "ed5a35dc-1b15-4487-af36-04374fe99235", // random UUID serving as magic string
		Sub: "Connection request",
		Exp: date,
		ID:  info.ID,
	}

	h, err := json.Marshal(headers)
	if err != nil {
		return "", fmt.Errorf("Error converting headers struct to JSON string: %s", err)
	}
	c, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("Error converting claims struct to JSON string: %s", err)
	}

	s := fmt.Sprintf("%s%s", h, c)
	return s, nil
}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func encodeJWT(s string) (string, error) {
	if IsJSON(s) {
		return "", errors.New("Error, JWT string is not valid JSON")
	}
	token, err := base64.URLEncoding.DecodeString(s)

	if err != nil {
		return "", fmt.Errorf("Base64 error: %s", err)
	}
	return string(token), nil
}
