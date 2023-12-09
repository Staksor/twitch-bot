package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"encoding/json"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// Prints a random interesting fact from API Ninjas
func Fact(
	message twitch.PrivateMessage,
	client *twitch.Client,
	cooldowns map[string]*time.Time,
) {
	iniData := utils.GetIniData()

	gptCooldown, _ := iniData.Section("cooldowns").Key("fact").Int()
	if !utils.CheckCooldown("fact", gptCooldown, message, client, cooldowns) {
		return
	}

	var apiResponse string = utils.ApiNinjasRequest("facts")

	var facts []structs.Fact

	json.Unmarshal([]byte(apiResponse), &facts)

	if len(facts) > 0 {
		client.Say(message.Channel, facts[0].Fact+" forsenScoots")
	}
}
