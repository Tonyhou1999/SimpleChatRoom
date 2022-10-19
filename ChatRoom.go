package SimpleChatRoom

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

var messageQueue chan Message
var connSlice map[string]net.Conn

func SendMessage() {
	var myMessage Message
	for {
		myMessage = <-messageQueue
		//todo make sure dont have to error check that myMessage.To will return non-nil
		conn, ok := connSlice[myMessage.To]
		if !ok {
			//send err
		}
		encoder := gob.NewEncoder(conn)
		err := encoder.Encode(myMessage)
		if err != nil {
			delete(connSlice, myMessage.To)
			//send err
		}
	}
}

// This file is used to create a chatroom that supports the messaging
func ClientThread(conn net.Conn) {
	var myMessage Message
	for {
		decoder := gob.NewDecoder(conn)
		decoder.Decode(&myMessage)
		messageQueue <- myMessage
	}
}

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
		go ClientThread(connection)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Incorrect args: should be ./[PROGNAME] port")
	}
	EstablishConnection(os.Args[1])
}
