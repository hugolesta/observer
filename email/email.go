package email

import "fmt"

type Email struct{}

func (e *Email) Notify(data string) {
	sendEmail(data)
}

func sendEmail(data string) {
	fmt.Printf("I've send an email with the message: %q", data)
}
