package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	PORT string = "8000"
)

// Start TCP server and accept connection requests
func main() {
	ln, err := net.Listen("tcp", "localhost:"+PORT)
	if err != nil {
		log.Fatalln("Server error:", err)
	}
	defer ln.Close()

	fmt.Printf("Server listening on http://localhost:%s\n\n", PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		go handleConnection(&conn)
	}
}

// Function to handle connection from client
func handleConnection(conn *net.Conn) {

	reader := bufio.NewReader(*conn)
	request, readErr := handleHTTPRequest(reader)
	if readErr != nil {
		log.Println("Failed to read request: ", readErr)
	}

	switch {
	case len(request["URL"]) >= 5 && request["URL"][:5] == "/user":
		userID := request["URL"][5:]
		err := userRequestHandler(userID, conn)
		if err != nil {
			log.Println("Cannot send data to client", err)
		}
	case request["URL"] == "/ws":
		if request["upgrade"] == "websocket" {
			err := upgradeToWebSocket(conn, request)
			if err != nil {
				log.Println("Write error, cannot upgrade to WebSocket:", err)
			}
			log.Println("WebSocket connection established")
			go readWSFrame(conn)
		} else {
			log.Println("Not a WebSocketUpgrade request")
			err := badRequestHandler(conn)
			if err != nil {
				log.Println("Cannot send Error 400 to client", err)
			}
		}

	default:
		err := notFoundHandler(conn)
		if err != nil {
			log.Println("Cannot send Error 404 to client", err)
		}
	}

}

// Function to handle HTTP request from client
func handleHTTPRequest(reader *bufio.Reader) (map[string]string, error) {
	request := make(map[string]string)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Println("line reading error", err, line)
	}
	log.Printf("Received:\n%s", line)
	parts := strings.Split(strings.TrimSpace(line), " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("Invalid request: %s", line)
	}

	request["Method"] = parts[0]
	request["URL"] = parts[1]
	request["Protocol"] = parts[2]

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			// end of request
			break
		} else if err != nil {
			log.Println("Error reading line:", line)
			break
		}
		fmt.Print(line)

		parts := strings.Split(strings.TrimSpace(line), ": ")
		if len(parts) == 1 {
			// end of request
			break
		} else if len(parts) != 2 {
			log.Println("Error in input, not 2 parts")
			break
		}
		request[parts[0]] = parts[1]
	}
	return request, err
}
