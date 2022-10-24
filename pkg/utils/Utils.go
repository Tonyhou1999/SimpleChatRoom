package utils

import "fmt"

type Message struct {
	To             string
	From           string
	MessageContent string
}

func (m Message) String() string {
	//todo error check for nil
	return fmt.Sprintf("To: %s\nFrom: %s\nMessage: %s", m.To, m.From, m.MessageContent)
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
