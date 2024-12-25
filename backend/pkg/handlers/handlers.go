package main

import (
	"fmt"
	"log"
	"net"
)

func userRequestHandler(userID string, conn *net.Conn) error {
	response := userID
	_, err := (*conn).Write([]byte(response))
	log.Printf("%s\r\n\r\n", response)
	return err
}

func badRequestHandler(conn *net.Conn) error {
	response := fmt.Sprintf("HTTP/1.1 400 NotFound\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nBad Request")
	_, err := (*conn).Write([]byte(response))
	log.Printf("%s\r\n\r\n", response)
	return err
}

func notFoundHandler(conn *net.Conn) error {
	response := fmt.Sprintf("HTTP/1.1 404 NotFound\r\nContent-Type: text/plain\r\nContent-Length: 9\r\n\r\nNot Found")
	_, err := (*conn).Write([]byte(response))
	log.Printf("%s\r\n\r\n", response)
	return err
}
