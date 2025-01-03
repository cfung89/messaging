package handlers

import (
	"fmt"
	"net"
)

func UserRequestHandler(userID string, conn *net.Conn) error {
	response := userID
	_, err := (*conn).Write([]byte(response))
	return err
}

func BadRequestHandler(conn *net.Conn) error {
	response := fmt.Sprintf("HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\nContent-Length: 11\r\n\r\nBad Request")
	_, err := (*conn).Write([]byte(response))
	return err
}

func NotFoundHandler(conn *net.Conn) error {
	response := fmt.Sprintf("HTTP/1.1 404 NotFound\r\nContent-Type: text/plain\r\nContent-Length: 9\r\n\r\nNot Found")
	_, err := (*conn).Write([]byte(response))
	return fmt.Errorf("Cannot send Error 404 to client: %s", err)
}

func TestHandler(conn *net.Conn) error {
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 7\r\n\r\nTest OK")
	_, err := (*conn).Write([]byte(response))
	return err
}

func UnauthorizedHandler(conn *net.Conn) error {
	response := fmt.Sprintf("HTTP/1.1 401 Unauthorized\r\nContent-Type: text/plain\r\nContent-Length: 13\r\n\r\nInvalid Login")
	_, err := (*conn).Write([]byte(response))
	return err
}
