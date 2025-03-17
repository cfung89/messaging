package main

import "github.com/cfung89/messaging/backend/pkg/auth"

func NewAuthServer(port int) *auth.AuthServer {
	s := &auth.AuthServer{}
	s.Server = s
	s.Start(8000)
	return s
}
