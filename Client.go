package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

/*
	connect to the TCP chatroom on client side,

and send the client's username tothe server
*/
func ConnectToChatRoom(port string, username string) {
	Port := ":" + port
	connection, err := net.Dial("tcp", Port)
	ConnectionError := "The provided Port Number is incorrect, please try again"
	check(err, ConnectionError) //Make sure it is a valid port number, i.e, the destination chatroom, exists
	UserInfo := Message{To: "chatroom", From: username, MessageContent: ""}
	encoder := gob.NewEncoder(connection)
	encoder.Encode(UserInfo)
	fmt.Printf("Current Client has been successfully registered, username:%s, port, %s", username, port)
	defer connection.Close()

}

// This function is designed to receive the actual message from the server, to the client
// It will print out the message received from the server
func receiveMessage(conn net.Conn, username string, message Message) {
	for {
		decoder := gob.NewDecoder(conn)
		err := decoder.Decode(message)
		check(err, "Server has closed")
		if err != nil {
			os.Exit(0)
		}
		fmt.Println("_______________")
		fmt.Printf("Message is from %s, following is the content\n", message.From)
		fmt.Printf("%s\n", message.MessageContent)
		fmt.Println("_______________")
	}
}
