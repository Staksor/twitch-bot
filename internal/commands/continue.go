package commands

import (
	"bot/internal/structs"

	"github.com/gempir/go-twitch-irc/v4"
)

// Continue stores GTP response for the last user request
func Continue(message twitch.PrivateMessage, client *twitch.Client, gptResponseStates map[string]*structs.GptResponseState) {
	userState := gptResponseStates[message.User.ID]

	if userState != nil && len(userState.Messages) > 0 {
		client.Reply(message.Channel, message.ID, userState.Messages[0])

		userState.Messages = userState.Messages[1:]

		if len(userState.Messages) > 0 {
			client.Reply(message.Channel, message.ID, "type !continue for more")
		}
		gptResponseStates[message.User.ID] = userState
	}
}
