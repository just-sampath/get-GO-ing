package main

import (
	"fmt"
	"sync"
)

type email struct {
	to, body string
}

func sendEmail(mail email, emailsChannel chan email) {
	fmt.Printf("Sending email to %s\n", mail.to)
	go func() {
		emailsChannel <- mail
		fmt.Printf("Sent an email to %s\n", mail.to)
		fmt.Println("===================================")
	}()
}

func receiveEmails(emailsChannel chan email) {
	for email := range emailsChannel {
		fmt.Printf("Received an email with body %s\n", email.body)
	}
}

func main() {
	emailChannels := make(chan email)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		receiveEmails(emailChannels)
		wg.Done()
	}()

	sendEmail(email{
		to:   "sampath@hello.com",
		body: "Hello!",
	}, emailChannels)

	close(emailChannels)
	wg.Wait()
}
