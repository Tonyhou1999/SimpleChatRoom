package utils

import "fmt"

type Message struct {
	To             string
	From           string
	MessageContent string
}

func Check(err error, message string) {
	if err != nil {
		fmt.Println(message)
	}
}

func CheckPanic(err error, message string) {
	if err != nil {
		panic(message)
	}
}
