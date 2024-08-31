package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	Connection *net.Conn
	Room       string
	Ping       chan bool // pinged client
	Timeout    *time.Ticker
	PingTimer  *time.Ticker
}

// Start client timers
func (client *Client) start() {
	for {
		select {
		case <-client.Timeout.C:
			log.Println("No pong")
			client.kill()
			return
		case <-client.PingTimer.C:
			err := sendPing(client.Connection)
			if err != nil {
				log.Println("Unable to send ping", err)
			}
			client.Timeout.Stop()
			client.Timeout = time.NewTicker(10 * time.Second)
			client.Ping <- true
		case val := <-client.Ping:
			if val == false {
				client.Timeout.Stop()
			}
		}
	}
}

// Close client
func (client *Client) kill() {
	close(client.Ping)
	(*client.Connection).Close()
	client.Timeout.Stop()
	client.PingTimer.Stop()

	room := client.Room
	client = nil // Client struct will be garbage collected from every room
	delete(chatrooms, room)
}

// Send ping to client
func sendPing(conn *net.Conn) error {
	frame := []byte{0x89, 0x0}
	_, err := (*conn).Write(frame)
	return err
}

// Send pong to client
func sendPong(conn *net.Conn, payload []byte) error {
	frame := []byte{0x8A}
	length := len(payload)
	if length <= 125 {
		frame = append(frame, byte(length))
	} else {
		return fmt.Errorf("Payload of ping frame too large")
	}

	frame = append(frame, payload...)

	_, err := (*conn).Write(frame)
	return fmt.Errorf("Unable to write pong frame to connection: %s", err)
}
