package main

import "encoding/json"

type MessageText struct {
	Author   string
	Room     string
	Type     string
	Text     string
	TimeSent *DateTime
}

// type MessageByte struct {
// 	Author   string
// 	Room     string
// 	Type     string
// 	Text     []byte
// 	TimeSent *DateTime
// }

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
