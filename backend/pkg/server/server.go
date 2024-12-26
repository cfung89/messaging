package server

import (
	"fmt"
	"log"
	"net"
)

type Server interface {
	Start(port int)                  // general for all (handles TCP connection)
	HandleConnection(conn *net.Conn) // specific for each server type
}

type BaseServer struct {
	server Server
}

func (b *BaseServer) Start(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println("Next port, taken: ", port)
		port += 1
	}
	defer ln.Close()
	log.Printf("Server listening on http://localhost:%d\n\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go b.server.HandleConnection(&conn)
	}
}
