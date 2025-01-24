package handlers

import (
	"fmt"
	"net"
)

func BadRequestHandler(conn *net.Conn) error {
	content := "Bad Request"
	response := fmt.Sprintf("HTTP/1.1 400 Bad Request\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len([]byte(content)), content)
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return err
}

func NotFoundHandler(conn *net.Conn) error {
	content := "Not Found"
	response := fmt.Sprintf("HTTP/1.1 404 NotFound\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len([]byte(content)), content)
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return fmt.Errorf("Cannot send Error 404 to client: %s", err)
}

func UnauthorizedHandler(conn *net.Conn) error {
	content := "Invalid login"
	response := fmt.Sprintf("HTTP/1.1 401 Unauthorized\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len([]byte(content)), content)
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return err
}

func InternalServerErrorHandler(conn *net.Conn) error {
	content := "Internal Server Error"
	response := fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len([]byte(content)), content)
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return err
}

func BadGatewayHandler(conn *net.Conn) error {
	content := "Bad Gateway"
	response := fmt.Sprintf("HTTP/1.1 502 Bad Gateway\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len([]byte(content)), content)
	_, err := (*conn).Write([]byte(response))
	(*conn).Close()
	return err
}
