package main

import ()

func main() {
	server := &Server{
		clients: make([]*Client, 0),
	}
	server.start()
}
