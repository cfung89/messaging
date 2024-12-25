package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Start client timers
func (client *Client) start() {
	for {
		select {
		case <-client.timeout.C:
			log.Println("No pong")
			client.kill()
			return
		case <-client.pingTimer.C:
			err := client.sendPing(client.connection)
			if err != nil {
				log.Println("Unable to send ping", err)
			}
			client.timeout.Stop()
			client.timeout = time.NewTicker(10 * time.Second)
			client.ping <- true
		case val := <-client.ping:
			if val == false {
				client.timeout.Stop()
			}
		}
	}
}

// Close client
func (client *Client) kill() {
	(*client.connection).Close()
	close(client.ping)
	client.timeout.Stop()
	client.pingTimer.Stop()

	room := client.room
	client = nil // Client struct will be garbage collected from every room
	delete(chatrooms, room)
}

// Send ping to client
func (client *Client) sendPing(conn *net.Conn) error {
	frame := []byte{0x89, 0x0}
	_, err := (*conn).Write(frame)
	return err
}

// Send pong to client
func (client *Client) sendPong(conn *net.Conn, payload []byte) error {
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
