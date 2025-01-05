package auth

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/cfung89/messaging/backend/pkg/assert"
	"github.com/cfung89/messaging/backend/pkg/db"
	"github.com/cfung89/messaging/backend/pkg/handlers"
	"github.com/cfung89/messaging/backend/pkg/server"
)

type AuthServer struct {
	server.BaseServer
	DB *sql.DB
}

type AuthObj struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AuthServer) InitDB() error {
	dbSecret, err := os.ReadFile("../secrets/authDB.key")
	if err != nil {
		return fmt.Errorf("Unable to read issuer from file: %s", err)
	}
	objs := strings.Split(string(dbSecret), "=")
	assert.Equal(&assert.EqualIn{A: objs[0], B: "AuthDbSecret", Err: "Invalid file read"})
	s.DB, err = db.OpenDB(objs[1], 5432)
	return err
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
