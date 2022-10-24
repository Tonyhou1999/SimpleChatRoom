package main

import (
	. "SimpleChatRoom/pkg/utils"
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

/*
	connect to the TCP chatroom on client side,

and send the client's username tothe server
*/

var encoder *gob.Encoder
var decoder *gob.Decoder

func UserInitialization() (chatroom string, username string) {
	if len(os.Args) != 3 {
		panic("The input is not correctly formatted. Type \"go run Client.go [chatroom number] [your desired username]\"")
	}
	//Now we need to check if the username is some reserved keywords. such as chatroom will be prohibited
	if strings.ToLower(os.Args[2]) == "chatroom" {
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

/*
This is used to initialize the message struct when the user is trying to send any information to the other client connected to server
Note from is evident so the user only needs to supply destination and message content, supplying EXIT at any, will cause disconnection to server, as suggested
*/
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

/*
The goal of this function is to provide the client to have access to the chatroom server, very similar to what we did for MP1
*/
func ConnectToChatRoom(address string, username string) net.Conn {
	connection, err := net.Dial("tcp", address)
	ConnectionError := "Not able to connect to the provided address, double check your address"
	CheckPanic(err, ConnectionError) //Make sure it is a valid port number, i.e, the destination chatroom, exists
	encoder = gob.NewEncoder(connection)
	decoder = gob.NewDecoder(connection)
	for {
		UserInfo := Message{To: "chatroom", From: username, MessageContent: ""}
		encoder.Encode(UserInfo)
		err := decoder.Decode(&UserInfo)
		Check(err, "Unable to receive response from server") //This is to make sure that the userinfo is received by the server

		//Detail on the username verification should check out ChatRoom.go file
		if err != nil {
			fmt.Println()
		}
		if UserInfo.MessageContent == "USERNAME ACCEPTED" {
			break
		} else {
			fmt.Println("Username was not accepted, please choose a different one")
			username = getInput("Username: ")
		}
	}
	fmt.Printf("Current Client has been successfully registered, username:%s, address:%s\n", username, address)
	return connection
}

func sendMessage(encoder *gob.Encoder, message Message) {
	err := encoder.Encode(message)
	Check(err, "Message was not sent to the server")
}

// This function is designed to receive the actual message from the server, to the client
// It will print out the message received from the server
func receiveMessage(conn net.Conn, username string) {
	var message Message
	for {
		err := decoder.Decode(&message)
		Check(err, "Server has closed")
		if err != nil && err != io.EOF {
			fmt.Println("Error here: ", err)
			os.Exit(0)
		}

		//This is the case to check if the client receives the EXIT sign from the server, it shall terminte shortly also
		if message.From == "chatroom" && message.MessageContent == "EXIT" {
			fmt.Println("Server will terminate the connection shortly")
			os.Exit(0)
		}
		fmt.Println("\n_______________")
		fmt.Printf("Message received from %s, following is the content\n", message.From)
		fmt.Printf("%s\n", message.MessageContent)
		fmt.Println("_______________")
	}
}

func main() {
	//Command Line arguments to gather client information
	port, username := UserInitialization()
	//send the client information to the actual server
	Connection := ConnectToChatRoom(port, username)
	go receiveMessage(Connection, username)
	fmt.Println("Input of each field of message should be read on separate lines." +
		" Enter \"EXIT \" to terminate the process")
	//The for loop will read the user input on each field of the message content
	//Will terminate in desired behavior
	for {
		message := MessageCreation(username)
		go sendMessage(encoder, message)
	}
}
