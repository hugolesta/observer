package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hugolesta/observer/email"
	"github.com/hugolesta/observer/message"
	"github.com/hugolesta/observer/observer"
	"github.com/hugolesta/observer/slack"
)

func main() {
	m := message.Message{}

	moreObservers := true

	for moreObservers {
		nameObs := readObserver()
		obs := observerFactory(nameObs)
		m.AddObserver(nameObs, obs)

		moreObservers = readAddMore()
	}

	m.Msg = readMessage()

	m.NotifyObservers()
}

func readMessage() string {
	fmt.Print("Enter message ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("It couldn't read the user message : %v", err)
	}
	text = strings.TrimSuffix(text, "\n")
	return text
}

func readObserver() string {
	fmt.Print("What observer do you need to add? ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("It couldn't read the user message : %v", err)
	}
	text = strings.TrimSuffix(text, "\n")
	return text
}

func readAddMore() bool {
	fmt.Println("Do youn need to add new observers? (y)/(n)")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalf("It couldn't read the user message : %v", err)
	}
	if char == 'y' {
		return true
	}
	return false
}

func observerFactory(name string) observer.Observer {
	switch name {
	case "slack":
		return &slack.Slack{}
	case "email":
		return &email.Email{}
	}
	return nil
}
