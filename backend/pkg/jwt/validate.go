package jwt

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cfung89/messaging/backend/pkg/assert"
	"github.com/cfung89/messaging/backend/pkg/utils"
)

func ValidateJWT(jwtStr string, filenames *SecretFilenames) (bool, error) {
	arr := strings.SplitN(jwtStr, ".", 3)
	assert.Equal(&assert.EqualIn{A: len(arr), B: 3, Err: "JWT is not 3 parts."})
	publicKey, err := utils.LoadPublicKey(filenames.PublicKey)
	if err != nil {
		return false, err
	}
	token := &Token{
		Raw:       jwtStr,
		Headers:   &Headers{},
		Claims:    &Claims{},
		Encoded:   [2]string{arr[0], arr[1]},
		Signature: &Signature{arr[2]},
		PublicKey: publicKey,
	}
	err = decodeJWT(token)
	if err != nil {
		return false, err
	}
	assert.NotNil(json.Unmarshal([]byte(token.Parts[0]), token.Headers))
	assert.NotNil(json.Unmarshal([]byte(token.Parts[1]), token.Claims))
	iss, err := os.ReadFile(filenames.Iss)
	if err != nil {
		return false, fmt.Errorf("Unable to read issuer from file: %s", err)
	}
	objs := strings.Split(string(iss), "=")
	assert.Equal(&assert.EqualIn{A: objs[0], B: "ISS_KEY", Err: "Invalid file read"})

	// Validation
	assert.Equal(&assert.EqualIn{A: *(token.Headers), B: Headers{Alg: "RS256", Typ: "JWT"}, Err: "Invalid JWT headers"})
	assert.Equal(&assert.EqualIn{A: token.Claims.Iss, B: objs[1], Err: "Invalid Issuer in JWT"})
	// assert.Equal(&assert.EqualIn{A: token.Claims.Sub, B: Subject, Err: "Invalid subject in JWT"})
	// assert.Equal(&assert.EqualIn{A: token.Claims.Username, B: Username, Err: "Invalid username in JWT"})
	duration := token.Claims.Exp - token.Claims.Iat
	assert.LessThan(&assert.LtIn{A: 0, B: duration, Err: "Invalid token times"})
	assert.LessThan(&assert.LtIn{A: 86400, B: duration, Err: "Expired token"})
	return validateSignature(token)
}

// Base64 URL encode JSON headers and JSON claims
func decodeJWT(token *Token) error {
	part0, err := base64.URLEncoding.DecodeString(token.Encoded[0])
	if err != nil {
		return fmt.Errorf("Error decoding string: %s", err)
	}
	part1, err := base64.URLEncoding.DecodeString(token.Encoded[1])
	if err != nil {
		return fmt.Errorf("Error decoding string: %s", err)
	}
	token.Parts = [2]string{string(part0), string(part1)}
	return nil
}

func validateSignature(token *Token) (bool, error) {
	hashed := sha256.Sum256([]byte(fmt.Sprintf("%s.%s.", token.Encoded[0], token.Encoded[1])))
	err := rsa.VerifyPKCS1v15(token.PublicKey, crypto.SHA256, hashed[:], []byte(token.Signature.string))
	if err != nil {
		return false, fmt.Errorf("Error verifying JWT signature: %s", err)
	}
	return true, nil
}
