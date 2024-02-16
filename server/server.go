// Server
package main

import (
	"fmt"
	"net"
)


func main() {

	
	// Listen for incoming connections on port 8080
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	// Accept incoming connection
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connection established with", conn.RemoteAddr())

	// Handle connection...
}
