package commands

import (
	"github.com/gempir/go-twitch-irc/v4"
)

func Test(message twitch.PrivateMessage, client *twitch.Client) {
	if message.User.Name == "staksor" {
		client.Say(message.Channel, "test")
	}
}
