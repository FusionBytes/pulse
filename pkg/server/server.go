package serverpackage

import (
	"fmt"
	"net"
	"os"
)

const (
	Host = "localhost"
	Port = "9988"
	Type = "tcp"
)

func main() {
	fmt.Println("Server running...")
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

	fmt.Println(fmt.Sprintf("Listenning on %s:%s", Host, Port))
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
			fmt.Println("Received: ", string(buffer[:mLen]))
			_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
			connection.Close()
		}(connection)
	}
}
