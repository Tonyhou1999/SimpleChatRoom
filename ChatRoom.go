package SimpleChatRoom

import (
	"fmt"
	"net"
)

//This file is used to create a chatroom that supports the messaging

// EstablishConnection refers to the TCP structure that starts building the TCP connection via the port
func EstablishConnection(InputPort string) {
	Port := ":" + InputPort
	listener, error := net.Listen("tcp", Port)
	ConnectionError := "The provided Port Number is incorrect, please try again"
	check(error, ConnectionError) //Here's the error checking part
	fmt.Println("Successful Establishment on Connecting Via Port" + Port + " is successful")
	defer listener.Close()
	for {
		connection, error := listener.Accept()
		check(error, "Please retry connection, there's an error here")
		fmt.Println(connection) //This is just for placeholder
	}
}
