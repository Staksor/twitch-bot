package commands

import (
	"bot/internal/utils"

	"github.com/gempir/go-twitch-irc/v4"
)

func JoinChannel(message twitch.PrivateMessage, client *twitch.Client, channel string) {
	iniData := utils.GetIniData()

	if message.Channel == iniData.Section("main").Key("main_channel").String() {
		client.Join(channel)
		client.Say(message.Channel, "Joined "+channel)
	}
}
