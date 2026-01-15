package main

import "fmt"

// 1. Notifier

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct{}

func (en EmailNotifier) Send(mail string) error {
	fmt.Println("Email: ", mail)
	return nil
}

type SMSNotifier struct{}

func (smsn SMSNotifier) Send(sms string) error {
	fmt.Println("SMS: ", sms)
	return nil
}

func SendNotification(n Notifier, msg string) {
	n.Send(msg)
}

// 2. Authentcator Interface
type Authenticator interface {
	AuthenticateUser(username, password string) bool
}

type SimpleAuthenticator struct {
	authenticator Authenticator
}

type AuthService struct{}

func (as AuthService) AuthenticateUser(username, password string) bool {
	return username == "admin" && password == "password"
}

func main() {
	sms := SMSNotifier{}
	email := EmailNotifier{}

	SendNotification(sms, "This notification is sent through SMS")
	SendNotification(email, "This Notification is sent through email")

	// 2. Authenticate User
	as := AuthService{}
	sa := SimpleAuthenticator{authenticator: as}

	fmt.Println(sa.authenticator.AuthenticateUser("admin", "password"))
}
