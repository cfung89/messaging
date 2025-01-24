package handlers

import (
	"fmt"
	"net"
)

func UserRequestHandler(userID string, conn *net.Conn) error {
	response := userID
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return err
}

func MessageSendHandler(conn *net.Conn, content string) error {
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len([]byte(content)), content)
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return err
}

func TestHandler(conn *net.Conn) error {
	content := "Test OK"
	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len([]byte(content)), content)
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return err
}
