package bots

import (
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// SlackRequest - request struct for the Slack API
type SlackRequest struct {
	Token   string
	Channel string
	Text    string
}

// SlackConfig - config for saving connection data and params
// for Slack connection
type SlackConfig struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

// SlackBot - simple bot for sending messages to Slack
type SlackBot struct {
	Config   SlackConfig `json:"config"`   // config for the channel
	Username string      `json:"username"` // username of the bot
}

// SendMessage - send message to the slack channel
func (bot SlackBot) SendMessage(text string) (*http.Response, error) {
	slackFields := &SlackRequest{
		Token:   bot.Config.Token,
		Channel: bot.Config.Channel,
		Text:    text,
	}

	slackFieldsBytes, convertErr := json.Marshal(slackFields)

	if convertErr != nil {
		return nil, convertErr
	}

	reader := strings.NewReader(string(slackFieldsBytes))

	request, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", reader)
	if err != nil {
		log.WithField("err", err.Error()).Error("Could not connect to the Slack channel.")
		return nil, err
	}
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.WithField("err", err.Error()).Error("Could not send message to Slack channel.")
		return resp, err
	}
	return resp, nil
}
