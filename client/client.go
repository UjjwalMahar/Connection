// Client
package main

import (
	"fmt"
	"net"
)

func main() {
	//data
	data := []byte("Hello")
	// Connect to server on port 8080
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	//try to write data
	conn.Write(data)

}
