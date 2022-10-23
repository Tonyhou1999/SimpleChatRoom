package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
