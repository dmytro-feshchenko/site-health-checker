package models

import (
	"net/smtp"
	"time"
)

// Notification - represents notification for the monitor
type Notification struct {
	Model

	Monitor  Monitor   `json:"monitor"`
	IsActive string    `json:"is_active"`
	LastUsed time.Time `json:"last_used"`
}

// Sender - interface for sending messages
type Sender interface {
	Send() error
}

// EmailNotification - represents email notification for the user
type EmailNotification struct {
	Notification

	Email string `json:"email" gorm:"email"`
}

// Send - send email notification to the client
// using SMTP
func (n EmailNotification) Send() error {
	auth := smtp.PlainAuth("", "feschenko.dimi3ryi@gmail.com", "password", "smtp.gmail.com")
	smtp.SendMail("smtp.gmail.com:587", auth, "feschenko.dimi3ryi@gmail.com", []string{"feschenko.dmitryi@gmail.com"}, []byte("test smtp server"))
	return nil
}

// SMSNotification - represents SMS notification for the user
type SMSNotification struct {
	Notification

	Number string `json:"number"`
}

// SlackNotification - represents notification for some Slack Channel
type SlackNotification struct {
	Notification

	Token   string `json:"token"`
	Channel string `json:"channel"`
}

// HipChatNotification - notification for HipChat
type HipChatNotification struct {
	Notification

	RoomID    string `json:"room_id"`
	AuthToken string `json:"auth_token"`
}
