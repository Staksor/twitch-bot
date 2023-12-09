package commands

import (
	"bot/internal/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// Makes a chat pyramid
func Pyramid(
	message twitch.PrivateMessage,
	client *twitch.Client,
	args []string,
	cooldowns map[string]*time.Time,
) {
	iniData := utils.GetIniData()

	cooldown, _ := iniData.Section("cooldowns").Key("pyramid").Int()
	if !utils.CheckCooldown("pyramid", cooldown, message, client, cooldowns) {
		return
	}

	var width int = 3
	var emote string = "TriHard"

	if len(args) > 0 {
		emote = args[0]
	}

	if len(args) > 1 {
		var err error
		width, err = strconv.Atoi(args[1])

		if err != nil {
			width = 3
		}
	}

	if width < 3 || width > 6 {
		width = 3
	}

	for i := 0; i < width; i++ {
		client.Say(message.Channel, strings.Repeat(emote+" ", i+1))
	}
	for i := width; i > 1; i-- {
		client.Say(message.Channel, strings.Repeat(emote+" ", i-1))
	}
}
