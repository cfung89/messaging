package rproxy

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"

	"github.com/cfung89/messaging/backend/pkg/assert"
	"github.com/cfung89/messaging/backend/pkg/jwt"
	"github.com/cfung89/messaging/backend/pkg/server"
)

const ServerPort = 8000

type ReverseProxy struct {
	server.BaseServer
	In  *net.Conn
	Out *net.Conn
}

func ForwardConnection(conn *net.Conn, request map[string]string) {
}

func (s *ReverseProxy) HandleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	request, readErr := s.BaseServer.HandleHTTPRequest(reader)
	if readErr != nil && readErr != io.EOF {
		log.Fatalf("Failed to read request: %s\n", readErr)
	}

	// JWT validation
	filenames := &jwt.SecretFilenames{
		Iss:       "../secrets/iss.env",
		PublicKey: "../secrets/public.pem",
	}
	authorization := strings.Split(request["Authorization"], " ")
	assert.Equal(&assert.EqualIn{A: authorization[0], B: "Bearer", Err: "Authorization header is not of type Bearer"})
	token := authorization[1]
	validJWT, err := jwt.ValidateJWT(token, filenames)
	if err != nil {
		log.Println(err)
	}
	if !validJWT {
		log.Println("Invalid JWT")
	}

	// forward to API server
}
