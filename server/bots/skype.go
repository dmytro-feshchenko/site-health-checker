package bots

import "net/http"

// SkypeConfig - contains configuration for the Skype API
type SkypeConfig struct {
}

// SkypeAuthenticationRequest - contains fields request
// for the authentication via Skype Bot Connector Server
type SkypeAuthenticationRequest struct {
	GrantType    string `json:"grant_type"`    // client credentials
	ClientID     string `json:"client_id"`     // Microsoft App ID
	ClientSecret string `json:"client_secret"` // Microsoft App Password
	Scope        string `json:"scope"`         // https://api.botframework.com/.default
}

// SkypeBot - bot for Skype messagner
type SkypeBot struct {
	Config          SkypeConfig `json:"skype_config"`
	IsAuthenticated bool        `json:"is_authenticated"`
}

// Authenticate - authenticate bot via Skype Bot Connector Server
// uses JWT authentication for the authorization
func (bot SkypeBot) Authenticate(clientID string, clientSecret string) error {
	// authenticationRequest := SkypeAuthenticationRequest{
	// 	GrantType:    "client_credentials",
	// 	ClientID:     clientID,
	// 	ClientSecret: clientSecret,
	// 	Scope:        "https://api.botframework.com/.default",
	// }
	// authenticationURL := "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"

	return nil
}

// SendMessage - send message to the
func (bot SkypeBot) SendMessage(text string) (*http.Response, error) {
	return nil, nil
}
