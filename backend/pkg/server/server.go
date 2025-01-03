package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server interface {
	Start(port int)                                                    // general for all (handles TCP connection)
	HandleConnection(conn *net.Conn)                                   // specific for each server type
	HandleHTTPRequest(reader *bufio.Reader) (map[string]string, error) // All servers use HTTP
}

type BaseServer struct {
	Server Server
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
		defer conn.Close()
		go b.Server.HandleConnection(&conn)
	}
}

// Function to handle HTTP request from client
func (b *BaseServer) HandleHTTPRequest(reader *bufio.Reader) (map[string]string, error) {
	request := make(map[string]string)
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		// end of request
		return nil, io.EOF
	} else if err != nil {
		log.Println("Error reading line:", line)
	}
	parts := strings.Split(strings.TrimSpace(line), " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("Invalid request: %s", line)
	}
	request["Method"] = parts[0]
	request["URL"] = parts[1]
	request["Protocol"] = parts[2]

	var contentLength int
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
		if parts[0] == "Content-Length" {
			contentLength, err = strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("Error converting Content-Length to integer: %s", err)
			}
		}
		request[parts[0]] = parts[1]
	}

	if contentLength > 0 {
		body := make([]byte, contentLength)
		_, err := io.ReadFull(reader, body)
		if err != nil {
			return nil, fmt.Errorf("Error reading body: %s", err)
		}
		request["Body"] = string(body)
	}

	log.Printf("Request: %s", request)
	return request, err
}
