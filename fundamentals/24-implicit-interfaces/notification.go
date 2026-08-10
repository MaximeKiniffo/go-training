package main

import "fmt"

type Notifiable interface {
	Notify() string
}

type Email struct {
	Address string
}

type SMS struct {
	Number string
}

func (e Email) Notify() string {
	return fmt.Sprintf("Notification par email envoyée à %s", e.Address)
}

func (s SMS) Notify() string {
	return fmt.Sprintf("Notification par SMS envoyée au numéro %s", s.Number)
}

func sendNotification(n Notifiable) {
	fmt.Println(n.Notify())
}

func main() {
	email := Email{Address: "john@example.com"}
	sms := SMS{Number: "1234567890"}
	sendNotification(email)
	sendNotification(sms)
}
