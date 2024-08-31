package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

// Upgrades connection to WebSocket
func upgradeToWebSocket(conn *net.Conn, request map[string]string) error {
	websocketAccept := generateWebSocketAccept(request["sec-websocket-key"])
	response := fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n", websocketAccept)

	fmt.Printf("Sending:\n%s\n", response)
	message := []byte(response)

	_, write_err := (*conn).Write(message)
	if write_err != nil {
		return write_err
	}
	return nil
}

// Generates WebSocket Accept key to conclude the handshake
func generateWebSocketAccept(key string) string {
	magicString := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	var websocketAccept string = fmt.Sprintf("%s%s", key, magicString)
	hasher := sha1.New()
	hasher.Write([]byte(websocketAccept))
	websocketAccept = base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return websocketAccept
}

// Reads all messages sent through WebSocket
func readWSFrame(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	decoded := make([]byte, 0)
	var client *Client

	for {
		firstByte, err := reader.ReadByte()
		if err != nil {
			log.Println("Error reading first data frame byte", err)
			if (&Client{}) == client {
				(*client).kill()
			} else {
				(*conn).Close()
			}
			return
		}

		fin := firstByte&0x80 != 0
		opcode := firstByte & 0x0F

		secondByte, err := reader.ReadByte()
		if err != nil {
			log.Println("Error reading length of payload", err)
		}

		temp := fmt.Sprintf("%b", secondByte)
		maskStr, err := strconv.ParseInt(string(temp[0]), 2, 64)
		if err != nil {
			log.Println("Error converting mask bit", err)
		}
		payloadLen, err := strconv.ParseInt(temp[1:], 2, 64)
		if err != nil {
			log.Println("Error converting payload length bits", err)
		}

		mask := maskStr == 1
		if mask == false {
			fmt.Println("Unmasked message from client")
			return
		}

		var length int

		switch {
		case payloadLen <= 125:
			length = int(payloadLen)
		case payloadLen == 126:
			payloadLen := make([]byte, 2)
			_, err := reader.Read(payloadLen)
			if err != nil {
				log.Println("Error reading length of payload", err)
			}
			length = int(binary.BigEndian.Uint16(payloadLen))
		case payloadLen == 127:
			payloadLen := make([]byte, 8)
			_, err := reader.Read(payloadLen)
			if err != nil {
				log.Println("Error reading length of payload", err)
			}
			length = int(binary.BigEndian.Uint16(payloadLen))
		default:
			log.Printf("Unsupported length: %d", payloadLen)
			// return fmt.Errorf("Unsupported length: %d", payloadLen)
		}

		maskKey := make([]byte, 4)
		_, err = reader.Read(maskKey)
		if err != nil {
			log.Println("Error reading mask key")
		}

		encoded := make([]byte, length)
		_, err = reader.Read(encoded)
		if err != nil {
			log.Println("Error reading encoded payload")
		}

		for i, n := range encoded {
			decoded = append(decoded, n^maskKey[i%4])
		}

		var message *MessageText
		switch opcode {
		case 0x0:
			// Continuation frame
			// pass
		case 0x1:
			// Text frame
			log.Printf("MESSAGE: %s\n", decoded)
			if fmt.Sprintf("%s", decoded) == "Hello Server" {
				err = sendWSFrame(conn, []byte("Hello Client"))
				if err != nil {
					log.Printf("Unable to write frame to client")
				}

				log.Println("Server Hello done")
				room := generateRoom() // placeholder function, will get room from client
				client = handleNewClient(conn, room)
				decoded = make([]byte, 0)
				continue
			}

			message, err = handleTextMessage(decoded) // Parse JSON message
			if err != nil {
				log.Println(err)
			}
		case 0x2:
			// Binary frame
			// pass
		case 0x8:
			// Connection close frame
			if (&Client{}) == client {
				(*client).kill()
			} else {
				(*conn).Close()
			}
			return
		case 0x9:
			// Ping frame
			err = sendPong(conn, decoded)
			if err != nil {
				log.Println(err)
			}
		case 0xA:
			// Pong frame
			if (&Client{}) == client {
				client.Ping <- false
			}
		default:
			log.Println("Error: Unknown opcode")
			if (&Client{}) == client {
				(*client).kill()
			} else {
				(*conn).Close()
			}
			return
		}

		if fin {
			chatroom := chatrooms[message.Room]
			for _, n := range chatroom {
				connect := *n.Connection
				connect.Write([]byte(decoded))
			}

			decoded = make([]byte, 0)
		}
	}
}

// Send WebSocket Frame to client
func sendWSFrame(conn *net.Conn, message []byte) error {
	frame := []byte{0x81}
	length := len(message)
	if length <= 125 {
		frame = append(frame, byte(length))
	} else if length <= 65535 {
		frame = append(frame, 126)
		lengthBytes := make([]byte, 2)
		binary.BigEndian.PutUint64(lengthBytes, uint64(length))
		frame = append(frame, lengthBytes...)
	} else {
		frame = append(frame, 127)
		lengthBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(lengthBytes, uint64(length))
		frame = append(frame, lengthBytes...)
	}
	frame = append(frame, message...)

	_, err := (*conn).Write(frame)
	return err
}

type MessageText struct {
	Author   string
	Room     string
	Type     string
	Text     string
	TimeSent *DateTime
}

type MessageByte struct {
	Author   string
	Room     string
	Type     string
	Text     []byte
	TimeSent *DateTime
}

type DateTime struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}

func handleTextMessage(decoded []byte) (*MessageText, error) {
	var message *MessageText
	err := json.Unmarshal([]byte(decoded), &message)
	return message, err
}
