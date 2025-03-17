package main

import (
	"github.com/cfung89/messaging/backend/pkg/auth"
)

func main() {
	s := &auth.AuthServer{}
	s.Server = s
	s.Start(8000)

	// r := &rproxy.ReverseProxy{}
	// r.Server = r
	// r.Start(8080)
}
