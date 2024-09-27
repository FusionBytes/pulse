package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	ServerHost = "localhost"
	ServerPort = "9988"
	ServerType = "tcp"
)

func main() {
	// Establish connection
	connection, err := net.Dial(ServerType, fmt.Sprintf("%s:%s", ServerHost, ServerPort))
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// Use a buffered reader to read user input
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		command = strings.TrimSpace(command)

		// Send command to the server
		_, err = connection.Write([]byte(command))
		if err != nil {
			fmt.Println("Error sending command:", err)
			return
		}

		if command == "exit" {
			fmt.Println("Exiting...")
			return
		}

		// Read response from the server
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Println(string(buffer[:mLen]))
	}
}
