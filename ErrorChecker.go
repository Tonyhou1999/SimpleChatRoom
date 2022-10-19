package SimpleChatRoom

import "fmt"

//Here an error checking method will be created to handle the error checking, as this code will always require to reuse
//As an modification, there will be messages to be printed, depending on where it is checked, a custom error message will display

func check(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}

func checkPanic(err error, message string) {
	if err != nil {
		panic(message)
	}
}
