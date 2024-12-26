package api

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/cfung89/messaging/backend/pkg/handlers"
	"github.com/cfung89/messaging/backend/pkg/server"
	"github.com/cfung89/messaging/backend/pkg/websocket"
)

type HTTPServer struct {
	server.BaseServer
}

// Function to handle connection from client
func (h *HTTPServer) HandleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	request, readErr := h.handleHTTPRequest(reader)
	if readErr != nil {
		log.Println("Failed to read request: ", readErr)
	}

	switch {
	case len(request["URL"]) >= 5 && request["URL"][:5] == "/user":
		// URL = "/user/USER_ID"
		userID := request["URL"][5:]
		err := handlers.UserRequestHandler(userID, conn)
		if err != nil {
			log.Println("Cannot send data to client", err)
		}
	case request["URL"] == "/ws":
		if request["upgrade"] == "websocket" {
			err := websocket.UpgradeToWebSocket(conn, request)
			if err != nil {
				log.Println("Write error, cannot upgrade to WebSocket:", err)
			}
			log.Println("WebSocket connection established")
			go websocket.ReadWSFrame(conn)
		} else {
			log.Println("Not a WebSocketUpgrade request")
			err := handlers.BadRequestHandler(conn)
			if err != nil {
				log.Println("Cannot send Error 400 to client", err)
			}
		}
	default:
		err := handlers.NotFoundHandler(conn)
		if err != nil {
			log.Println("Cannot send Error 404 to client", err)
		}
	}

}

// Function to handle HTTP request from client
func (h *HTTPServer) handleHTTPRequest(reader *bufio.Reader) (map[string]string, error) {
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
