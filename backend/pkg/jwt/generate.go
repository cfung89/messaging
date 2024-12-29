package jwt

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cfung89/messaging/backend/pkg/assert"
	"github.com/cfung89/messaging/backend/pkg/utils"
)

func GenerateToken(info *UserInfo) (*Token, error) {
	headers := &JWTHeaders{
		Alg: "RS256", // RSA SHA-256
		Typ: "JWT",
	}
	iss, err := os.ReadFile("./iss.env")
	if err != nil {
		return nil, fmt.Errorf("Unable to read issuer from file: %s", err)
	}
	objs := strings.Split(string(iss), "=")
	assert.Equal(&assert.EqualIn{A: objs[0], B: "ISS_KEY", Err: "Invalid file read"})
	claims := &JWTBody{
		Iss: objs[1], // magic string
		Sub: "Connection request",
		Exp: time.Now().AddDate(0, 0, 1).Unix(),
		ID:  info.ID,
	}
	token := &Token{
		Raw:       "",
		Header:    headers,
		Claims:    claims,
		Signature: nil,
	}

	h, c, err := convertJson(headers, claims)
	if err != nil {
		return nil, err
	}
	encodedH, err := encodeJWT(h)
	if err != nil {
		return nil, err
	}
	encodedC, err := encodeJWT(c)
	if err != nil {
		return nil, err
	}
	encoded := fmt.Sprintf("%s.%s.", encodedH, encodedC)
	token.Signature, err = generateSignature(encoded)
	if err != nil {
		return nil, err
	}
	token.Raw = fmt.Sprintf("%s%s", encoded, token.Signature.string)
	return token, nil
}

// Convert headers and claims to JSON string
func convertJson(headers *JWTHeaders, claims *JWTBody) (string, string, error) {
	h, err := json.Marshal(headers)
	if err != nil {
		return "", "", fmt.Errorf("Error converting headers struct to JSON string: %s", err)
	}
	c, err := json.Marshal(claims)
	if err != nil {
		return "", "", fmt.Errorf("Error converting claims struct to JSON string: %s", err)
	}
	return string(h), string(c), nil
}

// Base64 URL encode JSON headers and JSON claims
func encodeJWT(s string) (string, error) {
	if utils.IsJSON(s) {
		return "", errors.New("Error, JWT string is not valid JSON")
	}
	token := base64.URLEncoding.EncodeToString([]byte(s))
	// if err != nil {
	// 	return "", fmt.Errorf("Base64 error: %s", err)
	// }
	return string(token), nil
}

// Generate signature based on the encoded JSON
func generateSignature(s string) (*JWTSignature, error) {
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	if err != nil {
		return nil, fmt.Errorf("Error generating a random secret: %s", err)
	}
	data := []byte(s)
	hmac := hmac.New(sha256.New, secret) // new HMAC with hash type and key
	hmac.Write([]byte(data))             // give data to HMAC
	dataHmac := hmac.Sum(nil)            // compute HMAC
	hmacHex := hex.EncodeToString(dataHmac)
	// secretHex := hex.EncodeToString(secret) // key
	signature := &JWTSignature{hmacHex}
	return signature, nil
}
