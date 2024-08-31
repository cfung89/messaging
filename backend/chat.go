package main

import (
	"net"
	"time"
)

// Chatrooms map: consists of room IDs as keys and an array of Client objects as values (list of everyone in a chat)
var chatrooms = make(map[string][]*Client)

// Initialize new client
func handleNewClient(conn *net.Conn, room string) *Client {
	client := &Client{
		Connection: conn,
		Room:       room,
		Ping:       make(chan bool),
		Timeout:    time.NewTicker(10 * time.Second),
		PingTimer:  time.NewTicker(30 * time.Second),
	}

	_, ok := chatrooms[room]
	if ok {
		chatrooms[room] = append(chatrooms[room], client)
	} else {
		chatrooms[room] = []*Client{client}
	}

	go client.start()

	return client
}

// placeholder func
func generateRoom() string {
	return ""
}
