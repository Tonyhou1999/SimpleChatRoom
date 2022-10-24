package main

import (
	. "SimpleChatRoom/pkg/utils"
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

/*
	connect to the TCP chatroom on client side,

and send the client's username tothe server
*/

func UserInitialization() (chatroom string, username string) {
	if len(os.Args) != 3 {
		panic("The input is not correctly formatted. Type \"go run main.go [chatroom number] [your desired username]\"")
	}
	//Now we need to check if the username is some reserved keywords. such as chatroom will be prohibited
	if strings.ToLower(os.Args[1]) == "chatroom" {
		panic("The word chatroom is reserved for server use and can not be any part of the username")
	}
	chatroom, username = os.Args[1], os.Args[2]
	return
}

// This will prompt the message, and asks the user to input the relevant field of a message
func getInput(field string) (obtained string) {
	fmt.Print(field)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	obtained = scanner.Text()
	return
}
func MessageCreation(username string) (message Message) {
	to := getInput("To: ")
	info := getInput("Message: ")
	if to == "EXIT" || info == "EXIT" {
		fmt.Println("Command EXIT is received, program will terminate shortly")
		os.Exit(0)
	}
	message = Message{to, username, info}
	return
}

func ConnectToChatRoom(port string, username string) {
	Port := ":" + port
	connection, err := net.Dial("tcp", Port)
	ConnectionError := "The provided Port Number is incorrect, please try again"
	Check(err, ConnectionError) //Make sure it is a valid port number, i.e, the destination chatroom, exists
	UserInfo := Message{To: "chatroom", From: username, MessageContent: ""}
	encoder := gob.NewEncoder(connection)
	encoder.Encode(UserInfo)
	fmt.Printf("Current Client has been successfully registered, username:%s, port, %s", username, port)
	defer connection.Close()

}

func sendMessage(conn net.Conn, message Message) {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(message)
	Check(err, "Message was not sent to the server")
}

// This function is designed to receive the actual message from the server, to the client
// It will print out the message received from the server
func receiveMessage(conn net.Conn, username string, message Message) {
	for {
		decoder := gob.NewDecoder(conn)
		err := decoder.Decode(message)
		Check(err, "Server has closed")
		if err != nil {
			os.Exit(0)
		}
		fmt.Println("_______________")
		fmt.Printf("Message is from %s, following is the content\n", message.From)
		fmt.Printf("%s\n", message.MessageContent)
		fmt.Println("_______________")
	}
}
