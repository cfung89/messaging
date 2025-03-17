package auth

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/cfung89/messaging/backend/pkg/assert"
	"github.com/cfung89/messaging/backend/pkg/db"
	"github.com/cfung89/messaging/backend/pkg/handlers"
	"github.com/cfung89/messaging/backend/pkg/jwt"
	"github.com/cfung89/messaging/backend/pkg/server"
)

type AuthServer struct {
	server.BaseServer
	DB *db.AuthDB
}

type signUpResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type loginResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (s *AuthServer) InitDB(port int) error {
	if s.DB != nil {
		return errors.New("Server already has a database")
	}
	dbSecret, err := os.ReadFile("../secrets/authDB.key")
	if err != nil {
		return fmt.Errorf("Unable to read issuer from file: %s", err)
	}
	objs := strings.Split(string(dbSecret), "=")
	assert.Assert(objs[0] == "AuthDbSecret", "Invalid file read")
	s.DB = &db.AuthDB{}
	err = s.DB.OpenDB(objs[1], port)
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
	body := &db.UsersAuth{}
	err := json.Unmarshal([]byte(request["Body"]), body)
	if err != nil {
		log.Println(err)
	}

	switch {
	case request["URL"] == "/signup":
		err := s.DB.QueryUser(body)
		if err != sql.ErrNoRows {
			if err == nil {
				handlers.UnauthorizedHandler(conn)
				log.Println("Duplicate username")
				return
			}
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
			return
		}
		err = s.DB.InsertUsers(body)
		if err != nil {
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
			return
		}
		content, err := json.Marshal(&signUpResponse{Id: body.UID, Username: body.Username})
		if err != nil {
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
			return
		}
		response := fmt.Sprintf("HTTP/1.1 202 Accepted\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(content), string(content))
		_, err = (*conn).Write([]byte(response))
		if err != nil {
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
			return
		}

	case request["URL"] == "/login":
		err = s.DB.QueryUser(body)
		if err != nil {
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
			return
		}
		if body.DBPassword != body.Password {
			handlers.UnauthorizedHandler(conn)
			return
		}
		filenames := &jwt.SecretFilenames{Iss: "../secrets/iss.env", PrivateKey: "../secrets/private.pem"}
		token, err := jwt.GenerateToken(&jwt.UserInfo{ID: body.UID, Username: body.Username}, filenames)
		if err != nil {
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
			return
		}

		// Response
		content, err := json.Marshal(&loginResponse{Id: body.UID, Username: body.Username, Token: token.Raw})
		if err != nil {
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
			return
		}
		response := fmt.Sprintf("HTTP/1.1 202 Accepted\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(content), string(content))
		_, err = (*conn).Write([]byte(response))
		if err != nil {
			handlers.InternalServerErrorHandler(conn)
			log.Println(err)
		}
	default:
		err := handlers.NotFoundHandler(conn)
		if err != nil {
			log.Println(err)
		}
	}
}
