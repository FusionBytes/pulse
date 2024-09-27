package main

import (
	"fmt"
	"hash/fnv"
	"net"
	"os"
	"pulse/internal/commands"
	stringcommands "pulse/internal/commands/string"
	"pulse/internal/parser"
	"pulse/pkg/structure"
)

const (
	Host = "localhost"
	Port = "9988"
	Type = "tcp"
)

func main() {
	fmt.Println("Server running...")
	scoreboards := structure.NewHashTable(fnv.New64a(), 8, 0.75)
	parser := parser.NewParser(
		commands.NewZAdd(scoreboards),
		commands.NewZRange(scoreboards),
		commands.NewZRank(scoreboards),
		commands.NewZScore(scoreboards),
		stringcommands.NewSET(),
		stringcommands.NewGET(),
		// Add more commands as needed
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
			defer connection.Close() // Ensure the connection is closed when done

			buffer := make([]byte, 1024)
			for {
				mLen, err := connection.Read(buffer)
				if err != nil {
					fmt.Println("Error reading:", err.Error())
					return
				}
				cmd := string(buffer[:mLen])

				result, err := parser.Execute(cmd)
				if err != nil {
					_, _ = connection.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
					continue
				}

				if result == nil {
					_, _ = connection.Write([]byte(fmt.Sprintf("%v", result)))
				} else if intResult, ok := result.(int); ok {
					_, _ = connection.Write([]byte(fmt.Sprintf("%d", intResult)))
				} else if strResult, ok := result.(string); ok {
					_, _ = connection.Write([]byte(fmt.Sprintf("%s", strResult)))
				} else if boolResult, ok := result.(bool); ok {
					_, _ = connection.Write([]byte(fmt.Sprintf("%v", boolResult)))
				}
			}
		}(connection)
	}
}
