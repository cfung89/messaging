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
	assert.Assert(len(arr) == 3, "JWT is not 3 parts")
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
	assert.Error(json.Unmarshal([]byte(token.Parts[0]), token.Headers))
	assert.Error(json.Unmarshal([]byte(token.Parts[1]), token.Claims))
	iss, err := os.ReadFile(filenames.Iss)
	if err != nil {
		return false, fmt.Errorf("Unable to read issuer from file: %s", err)
	}
	objs := strings.Split(string(iss), "=")
	assert.Assert(objs[0] == "ISS_KEY", "Invalid file read")

	// Validation
	assert.Assert(token.Headers.Alg == "RS256", "Invalid JWT headers Alg")
	assert.Assert(token.Headers.Typ == "JWT", "Invalid JWT headers Typ")
	assert.Assert(token.Claims.Iss == objs[1], "Invalid Issuer in JWT")
	// assert.Assert(token.Claims.Sub == Subject, "Invalid subject in JWT"})
	// assert.Assert(token.Claims.Username == Username, "Invalid username in JWT"})
	duration := token.Claims.Exp - token.Claims.Iat
	assert.Assert(0 < duration, "Invalid token times")
	assert.Assert(86400 > duration, "Expired token")
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
