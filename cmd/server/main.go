package main

import (
	"fmt"
	"net"
	"os"
	"pulse/internal/commands/mock"
	"pulse/internal/parser"
)

const (
	Host = "localhost"
	Port = "9988"
	Type = "tcp"
)

func main() {
	fmt.Println("Server running...")
	parser := parser.NewParser(
		&mock.MockCommand{},
	)
	server, err := net.Listen(Type, fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			fmt.Println("Error closing:", err.Error())
			os.Exit(1)
		}
	}(server)

	fmt.Printf("Listening on %s:%s\n", Host, Port)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go func(connection net.Conn) {
			buffer := make([]byte, 1024)
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			cmd := string(buffer[:mLen])

			result, err := parser.Execute(cmd)

			if err != nil {
				_, _ = connection.Write([]byte(fmt.Sprintf("there is an error %s", err.Error())))
				connection.Close()
				return
			}

			if intResult, ok := result.(int); ok {
				_, _ = connection.Write([]byte(fmt.Sprintf("the result is %d", intResult)))
				connection.Close()
				return
			}


			if strResult, ok := result.(string); ok {
				_, _ = connection.Write([]byte(fmt.Sprintf("the result is %s", strResult)))
				connection.Close()
				return
			}
		}(connection)
	}
}
