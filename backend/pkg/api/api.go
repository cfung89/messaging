package api

import (
	"bufio"
	"log"
	"net"

	"github.com/cfung89/messaging/backend/pkg/handlers"
	"github.com/cfung89/messaging/backend/pkg/server"
	"github.com/cfung89/messaging/backend/pkg/websocket"
)

type AppServer struct {
	*server.BaseServer
}

// Function to handle connection from client
func (s *AppServer) HandleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	request, readErr := s.BaseServer.HandleHTTPRequest(reader)
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
			log.Println(err)
		}
	}

}
