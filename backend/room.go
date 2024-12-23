package main

import (
	"net"
)

type Room struct {
	ID    string
	Users map[string]*Client
}

// Chatrooms map: consists of room IDs as keys and an array of Client objects as values (list of everyone in a chat)
var chatrooms = make(map[string][]*Client)

// Initialize new client
func handleNewClient(conn *net.Conn, room string) *Client {
	client := &Client{
		connection: conn,
		room:       room,
		// ping:       make(chan bool),
		// timeout:    time.NewTicker(10 * time.Second),
		// pingTimer:  time.NewTicker(30 * time.Second),
	}

	_, ok := chatrooms[room]
	if ok {
		chatrooms[room] = append(chatrooms[room], client)
	} else {
		chatrooms[room] = []*Client{client}
	}

	// go client.start()

	return client
}

// placeholder func
func generateRoom() string {
	return ""
}
