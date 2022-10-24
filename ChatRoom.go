package main

// This file is used to create a chatroom that supports the messaging

import (
	. "SimpleChatRoom/pkg/utils"
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

var messageQueue chan Message
var mutex sync.Mutex

// todo resize connSlice when adding stuff
var connSlice map[string]net.Conn = make(map[string]net.Conn, 10)
var encodeSlice map[string]*gob.Encoder = make(map[string]*gob.Encoder, 10) //so don't create a new encoder each time

func sendErr(myMessage Message) {
	conn, ok := connSlice[myMessage.From]
	if !ok {
		fmt.Println("Unable to send error message back to " + myMessage.From)
	}
	//still using gob so can perhaps detect an error
	encoder, ok := encodeSlice[myMessage.To]
	if !ok {
		encoder = gob.NewEncoder(conn)
		encodeSlice[myMessage.To] = encoder
	}
	err := encoder.Encode("Unable to send message to " + myMessage.To)
	if err != nil {
		fmt.Println("Unable to send error message back to " + myMessage.From)
	}
}

func SendMessage() {
	var myMessage Message
	for {
		myMessage = <-messageQueue
		//todo make sure dont have to error check that myMessage.To will return non-nil
		conn, ok := connSlice[myMessage.To]
		if !ok {
			sendErr(myMessage)
			continue
		} else {
			encoder, ok := encodeSlice[myMessage.To]
			if !ok {
				encoder = gob.NewEncoder(conn)
				encodeSlice[myMessage.To] = encoder
			}
			err := encoder.Encode(myMessage)
			if err != nil {
				mutex.Lock()
				delete(connSlice, myMessage.To)
				delete(encodeSlice, myMessage.To)
				mutex.Unlock()
				sendErr(myMessage)
				continue
			}
		}
	}
}

func getUsername(conn net.Conn, decoder *gob.Decoder) {
	var myMessage Message
	myEncoder := gob.NewEncoder(conn)
	for {
		err := decoder.Decode(&myMessage)
		if err != nil {
			errMessage := Message{myMessage.From, "chatroom",
				"Unable to parse message, please try again"}
			myEncoder.Encode(errMessage)
		}
		if myMessage.To != "chatroom" {
			errMessage := Message{myMessage.From, "chatroom",
				"Please input a username, by sending a message to \"chatroom\" with a unique username"}
			myEncoder.Encode(errMessage)
		}
		_, hasVal := encodeSlice[myMessage.To]
		if hasVal {
			errMessage := Message{myMessage.From, "chatroom",
				"The username:\"" + myMessage.From + "\"is already taken, try another username"}
			myEncoder.Encode(errMessage)
		}
		mutex.Lock()
		encodeSlice[myMessage.From] = myEncoder
		connSlice[myMessage.From] = conn
		mutex.Unlock()
		return
	}
}

func ClientThread(conn net.Conn) {
	var myMessage Message
	decoder := gob.NewDecoder(conn)
	getUsername(conn, decoder)
	for {
		err := decoder.Decode(&myMessage)
		if err != nil && err != io.EOF {
			myMessage = Message{} //since we don't know where the message is From
			sendErr(myMessage)
			continue
		}
		messageQueue <- myMessage
	}
}

// establishConnection refers to the TCP structure that starts building the TCP connection via the port
func establishConnection(InputPort string) {
	Port := "127.0.0.1:" + InputPort
	listener, err := net.Listen("tcp4", Port)
	ConnectionError := "The provided Port Number is incorrect, exiting"
	CheckPanic(err, ConnectionError) //Here's the error checking part
	fmt.Println("Successfully started listening at " + listener.Addr().String())
	go SendMessage()
	defer listener.Close()
	for {
		connection, error := listener.Accept()
		Check(error, "Please retry connection, there's an error here")
		fmt.Println(connection) //This is just for placeholder
		go ClientThread(connection)
	}
}

func waitForEnd() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Type \"EXIT\" to close the chatroom")
		scanner.Scan()
		if scanner.Text() == "EXIT" {
			fmt.Println("Chatroom will be terminated shortly")
			os.Exit(0)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Incorrect args: should be ./[PROGNAME] PORT")
	}
	go establishConnection(os.Args[1])
	waitForEnd()
}
