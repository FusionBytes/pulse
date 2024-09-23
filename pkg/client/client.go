package client

import (
	"fmt"
	"net"
)

const (
	ServerHost = "localhost"
	ServerPort = "9988"
	ServerType = "tcp"
)

func main() {
	//establish connection
	connection, err := net.Dial(ServerType, fmt.Sprintf("%s:%s", ServerHost, ServerPort))
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	///send some data
	_, err = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
}
