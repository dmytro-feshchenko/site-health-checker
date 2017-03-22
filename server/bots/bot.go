package bots

import "net/http"

// Bot - interface for bots for sending messages
type Bot interface {
	SendMessage(text string) (*http.Response, error)
}
