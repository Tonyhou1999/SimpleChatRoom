package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
   This code is designed to handle the setting and termination of the connection
*/

// Getting the designed port in order for a successful establishment of ChatRoom, used for the EstablishConnection in ChatRoom.go file
func initialization(Port string) string {
	if len(os.Args) != 2 {
		panic("The provided input port message is not correct, enter in the format \"go run main.go Port \" ")
	}
	chatport := os.Args[1]
	return chatport
	//Whatever value returned here will be used for the EstablishConnection function, but it keeps reporting an issue
}

// The method below will be to inform the user on how to terminate the chatroom, which is by typing "EXIT"
func termination() {
	for {
		fmt.Println("Type \"EXIT\" to close the chatroom")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "EXIT" {
			fmt.Println("Chatroom will be terminated shortly")
			os.Exit(0)
		}
	}

}
