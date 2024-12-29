package jwt

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/cfung89/messaging/backend/pkg/assert"
)

func ValidateJWT(jwtStr string, filenames *SecretFilenames) (bool, error) {
	arr := strings.Split(jwtStr, ".")
	assert.Equal(&assert.EqualIn{A: len(arr), B: 3, Err: "JWT is not 3 parts."})
	return true, nil
}

func validateSignature(data string, token *Token) (bool, error) {
	hashed := sha256.Sum256([]byte(data))
	err := rsa.VerifyPKCS1v15(token.PublicKey, crypto.SHA256, hashed[:], []byte(token.Signature.string))
	if err != nil {
		return false, fmt.Errorf("Error verifying JWT signature: %s", err)
	}
	return true, nil
}
