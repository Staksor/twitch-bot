package commands

import (
	"bot/internal/utils"

	"github.com/gempir/go-twitch-irc/v4"
)

func LeaveChannel(message twitch.PrivateMessage, client *twitch.Client, channel string) {
	iniData := utils.GetIniData()

	if message.Channel == iniData.Section("main").Key("main_channel").String() {
		client.Depart(channel)
		client.Say(message.Channel, "Left "+channel)
	}
}
