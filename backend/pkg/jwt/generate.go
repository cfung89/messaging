package jwt

import (
	"crypto"
	"crypto/rand"
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

func GenerateToken(info *UserInfo, filenames *SecretFilenames) (*Token, error) {
	iss, err := os.ReadFile(filenames.Iss)
	if err != nil {
		return nil, fmt.Errorf("Unable to read issuer from file: %s", err)
	}
	objs := strings.Split(string(iss), "=")
	assert.Assert(objs[0] == "ISS_KEY", "Invalid file read")
	privateKey, err := utils.LoadPrivateKey(filenames.PrivateKey)
	if err != nil {
		return nil, err
	}
	token := &Token{
		Headers: &Headers{Alg: "RS256", Typ: "JWT"},
		Claims: &Claims{
			Iss:      objs[1], // magic string
			Sub:      info.ID,
			Username: info.Username,
			Iat:      time.Now().Unix(),
			Exp:      time.Now().AddDate(0, 0, 1).Unix(),
		},
		Parts:      [2]string{"", ""},
		Encoded:    [2]string{"", ""},
		PrivateKey: privateKey,
	}
	err = convertJson(token)
	if err != nil {
		return nil, err
	}
	err = encodeJWT(token)
	if err != nil {
		return nil, err
	}
	err = generateSignature(token)
	if err != nil {
		return nil, err
	}
	token.Raw = fmt.Sprintf("%s.%s.%s", token.Encoded[0], token.Encoded[1], token.Signature.string)
	return token, nil
}

// Convert headers and claims to JSON string
func convertJson(token *Token) error {
	h, err := json.Marshal(token.Headers)
	if err != nil {
		return fmt.Errorf("Error converting headers struct to JSON string: %s", err)
	}
	c, err := json.Marshal(token.Claims)
	if err != nil {
		return fmt.Errorf("Error converting claims struct to JSON string: %s", err)
	}
	token.Parts[0], token.Parts[1] = string(h), string(c)
	return nil
}

// Base64 URL encode JSON headers and claims
func encodeJWT(token *Token) error {
	if !utils.IsJSON(token.Parts[0]) || !utils.IsJSON(token.Parts[1]) {
		return errors.New("Error, JWT string is not valid JSON")
	}
	token.Encoded[0] = base64.URLEncoding.EncodeToString([]byte(token.Parts[0]))
	token.Encoded[1] = base64.URLEncoding.EncodeToString([]byte(token.Parts[1]))
	return nil
}

// Generate signature based on the encoded JSON
func generateSignature(token *Token) error {
	hashed := sha256.Sum256([]byte(fmt.Sprintf("%s.%s.", token.Encoded[0], token.Encoded[1])))
	signature, err := rsa.SignPKCS1v15(rand.Reader, token.PrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return fmt.Errorf("Error generating JWT signature: %s", err)
	}
	token.Signature = &Signature{string(signature)}
	return nil
}
