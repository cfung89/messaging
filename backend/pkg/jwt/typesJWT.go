package jwt

type Token struct {
	Raw       string        // encoded
	Header    *JWTHeaders   // decoded
	Claims    *JWTBody      // decoded
	Signature *JWTSignature // decoded
}

type JWTHeaders struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
	Cty string `json:"cty"`
}

type JWTBody struct {
	Iss  string `json:"iss"`  // issuer
	Sub  string `json:"sub"`  // subject
	Exp  int64  `json:"exp"`  // expiration time (1 day)
	Name string `json:"name"` // Name of user
	ID   string `json:"id"`   // ID of user
}

type JWTSignature struct {
	string
}

type UserInfo struct {
	Email    string
	Password string
	ID       string
}
