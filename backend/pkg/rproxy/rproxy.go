package rproxy

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"

	"github.com/cfung89/messaging/backend/pkg/assert"
	"github.com/cfung89/messaging/backend/pkg/handlers"
	"github.com/cfung89/messaging/backend/pkg/jwt"
	"github.com/cfung89/messaging/backend/pkg/server"
)

const ServerPort = 8000

type ReverseProxy struct {
	server.BaseServer
	Client *net.Conn
	Server *net.Conn
}

func NewRProxy(serverConn *net.Conn) ReverseProxy {
	return ReverseProxy{
		Client: nil,
		Server: serverConn,
	}
}

func (s *ReverseProxy) ForwardConnection() {
	reader := bufio.NewReader(*s.Client)
	request, err := s.BaseServer.HandleHTTPRequest(reader)
	if err != nil && err != io.EOF {
		log.Fatalf("Failed to read request: %s\n", err)
	}
	if request[""] == "Server" {
		err = handlers.MessageSendHandler(s.Client, request["Body"])
		if err != nil {
			log.Println(err)
		}
	} else if request[""] == "Client" {
		err := handlers.MessageSendHandler(s.Server, request["Body"])
		if err != nil {
			log.Println(err)
		}
	}
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
	assert.Assert(authorization[0] == "Bearer", "Authorization header is not of type Bearer")
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
