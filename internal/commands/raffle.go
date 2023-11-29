package commands

import (
	"bot/internal/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// Prints a random interesting fact from API Ninjas
func Raffle(message twitch.PrivateMessage, client *twitch.Client, seconds string) {
	iniData := utils.GetIniData()

	if message.User.Name == iniData.Section("main").Key("main_channel").String() {
		seconds, err := strconv.Atoi(seconds)

		if err == nil {
			client.Say(message.Channel, fmt.Sprintf("!raffle 100k %d", seconds))
			time.Sleep(1 * time.Second)
			client.Say(message.Channel, "!join")
		}
	}
}
