package jwt_test

import (
	"log"
	"testing"

	"github.com/cfung89/messaging/backend/pkg/jwt"
)

func TestGenerate(t *testing.T) {
	userinfo := &jwt.UserInfo{}
	filenames := &jwt.SecretFilenames{
		Iss:        "../secrets/iss.env",
		PrivateKey: "../secrets/private.pem",
		PublicKey:  "../secrets/public.pem",
	}
	token, err := jwt.GenerateToken(userinfo, filenames)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(token.Raw)
}

func TestValidation(t *testing.T) {
	userinfo := &jwt.UserInfo{}
	filenames := &jwt.SecretFilenames{
		Iss:        "../secrets/iss.env",
		PrivateKey: "../secrets/private.pem",
		PublicKey:  "../secrets/public.pem",
	}
	token, err := jwt.GenerateToken(userinfo, filenames)
	if err != nil {
		log.Fatalln(err)
	}
	valid, err := jwt.ValidateJWT(token.Raw, filenames)
	if err != nil {
		log.Fatalln(err)
	}
	if !valid {
		log.Fatalln("invalid JWT")
	}
}
