package utils

import (
	"fmt"
	"math"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

func CheckCooldown(
	cooldownType string,
	cooldown int,
	message twitch.PrivateMessage,
	client *twitch.Client,
	cooldowns map[string]*time.Time,
) bool {
	iniData := GetIniData()

	if iniData.Section("main").Key("main_channel").String() == message.User.Name {
		return true
	}

	timezone, err := time.LoadLocation("Europe/Berlin")

	if err != nil {
		fmt.Println("Error loading location:", err)
		return false
	}

	time.Local = timezone
	now := time.Now()

	cooldownByType := cooldowns[cooldownType]
	if cooldownByType == nil || now.After(*cooldownByType) {
		newTimestamp := time.Now().Add(time.Second * time.Duration(cooldown))
		cooldowns[cooldownType] = &newTimestamp

		return true
	} else {
		timeLeft := now.Sub(*cooldownByType)
		secondsLeft := int64(math.Abs(timeLeft.Seconds()))
		client.Reply(message.Channel, message.ID, fmt.Sprintf("Cooldown Sadeg (%ds left)", secondsLeft))

		return false
	}
}
