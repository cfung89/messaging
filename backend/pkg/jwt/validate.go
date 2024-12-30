package jwt

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cfung89/messaging/backend/pkg/assert"
	"github.com/cfung89/messaging/backend/pkg/utils"
)

func ValidateJWT(jwtStr string, filenames *SecretFilenames) (bool, error) {
	arr := strings.Split(jwtStr, ".")
	assert.Equal(&assert.EqualIn{A: len(arr), B: 3, Err: "JWT is not 3 parts."})
	decodedH, err := decodeJWT(arr[0])
	if err != nil {
		return false, err
	}
	decodedC, err := decodeJWT(arr[1])
	if err != nil {
		return false, err
	}
	var (
		headers *Headers
		claims  *Claims
	)
	json.Unmarshal([]byte(decodedH), headers)
	json.Unmarshal([]byte(decodedC), claims)
	assert.Equal(&assert.EqualIn{A: headers, B: &Headers{Alg: "RS256", Typ: "JWT"}, Err: "Invalid JWT headers"})
	iss, err := os.ReadFile(filenames.Iss)
	if err != nil {
		return false, fmt.Errorf("Unable to read issuer from file: %s", err)
	}
	objs := strings.Split(string(iss), "=")
	assert.Equal(&assert.EqualIn{A: objs[0], B: "ISS_KEY", Err: "Invalid file read"})
	assert.Equal(&assert.EqualIn{A: claims.Iss, B: iss, Err: "Invalid Issuer in JWT"})
	assert.Equal(&assert.EqualIn{A: claims.Sub, B: "Connection request", Err: "Invalid subject in JWT"})
	assert.LessThan(&assert.LtIn{A: claims.Exp, B: time.Now().Unix(), Err: "Expired token"})
	publicKey, err := utils.LoadPublicKey(filenames.PublicKey)
	if err != nil {
		return false, err
	}
	token := &Token{
		Raw:       jwtStr,
		Header:    headers,
		Claims:    claims,
		Signature: &Signature{arr[2]},
		PublicKey: publicKey,
	}
	return validateSignature(token)
}

// Base64 URL encode JSON headers and JSON claims
func decodeJWT(s string) (string, error) {
	if utils.IsJSON(s) {
		return "", errors.New("Error, JWT string is not valid JSON")
	}
	token, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("Error decoding string: %s", err)
	}
	return string(token), nil
}

func validateSignature(token *Token) (bool, error) {
	hashed := sha256.Sum256([]byte(token.Signature.string))
	err := rsa.VerifyPKCS1v15(token.PublicKey, crypto.SHA256, hashed[:], []byte(token.Signature.string))
	if err != nil {
		return false, fmt.Errorf("Error verifying JWT signature: %s", err)
	}
	return true, nil
}
