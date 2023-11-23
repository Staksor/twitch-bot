package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// Continues stored GTP response for the last user request
func Continue(
	message twitch.PrivateMessage,
	client *twitch.Client,
	gptResponseStates map[string]*structs.GptResponseState,
	cooldowns map[string]*time.Time,
) {
	userState := gptResponseStates[message.User.ID]
	iniData := utils.GetIniData()

	continueCooldown, _ := iniData.Section("cooldowns").Key("continue").Int()
	if !utils.CheckCooldown("continue", continueCooldown, message, client, cooldowns) {
		return
	}

	if userState != nil && len(userState.Messages) > 0 {
		client.Reply(message.Channel, message.ID, "🤖 "+userState.Messages[0])

		userState.Messages = userState.Messages[1:]

		if len(userState.Messages) > 0 {
			client.Reply(message.Channel, message.ID, "type !continue for more")
		}
		gptResponseStates[message.User.ID] = userState
	}
}
