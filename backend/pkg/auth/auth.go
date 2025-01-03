package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/cfung89/messaging/backend/pkg/handlers"
	"github.com/cfung89/messaging/backend/pkg/server"
)

type AuthServer struct {
	server.BaseServer
}

type AuthObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AuthServer) HandleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	request, readErr := s.BaseServer.HandleHTTPRequest(reader)
	if readErr != nil && readErr != io.EOF {
		log.Fatalf("Failed to read request: %s\n", readErr)
	}
	if request["Method"] != "POST" {
		err := handlers.BadRequestHandler(conn)
		if err != nil {
			log.Fatalln(err)
		}
	}
	body := &AuthObj{}
	err := json.Unmarshal([]byte(request["Body"]), body)
	if err != nil {
		log.Println(err)
	}

	switch {
	case request["URL"] == "/signup":
		log.Println("signup detected")

		err := handlers.TestHandler(conn)
		if err != nil {
			log.Println(err)
		}
		(*conn).Close()
	case request["URL"] == "/login":
		log.Println("login detected")

		// check database for username/password
		// if username/password not valid: handlers.UnauthorizedHandler()
		// return
		// if valid
		var content string
		contentLength := len([]byte(content))
		response := fmt.Sprintf("HTTP/1.1 202 Accepted\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", contentLength, content)
		_, err := (*conn).Write([]byte(response))

		err = handlers.TestHandler(conn)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("response sent")
		(*conn).Close()
	default:
		err := handlers.NotFoundHandler(conn)
		if err != nil {
			log.Println(err)
		}
		(*conn).Close()
	}
	return
}
