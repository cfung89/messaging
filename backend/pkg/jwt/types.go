package jwt

import "crypto/rsa"

type Token struct {
	Raw        string     // encoded
	Headers    *Headers   // decoded
	Claims     *Claims    // decoded
	Signature  *Signature // decoded
	Parts      [2]string  // not encoded
	Encoded    [2]string  // encoded string (both parts) without signature
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

type Headers struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Claims struct {
	Iss      string `json:"iss"`      // issuer
	Sub      string `json:"sub"`      // ID of user
	Iat      int64  `json:"iat"`      // issued at time
	Exp      int64  `json:"exp"`      // expiration time (1 day after IAT)
	Username string `json:"username"` // Name of user
}

type Signature struct {
	string
}

type UserInfo struct {
	ID       string
	Username string // not unique
	Email    string
	Password string
}

type SecretFilenames struct {
	Iss        string
	PrivateKey string
	PublicKey  string
}
