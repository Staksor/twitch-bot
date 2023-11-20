package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"encoding/json"

	"github.com/gempir/go-twitch-irc/v4"
)

// Prints a random interesting fact from API Ninjas
func Fact(message twitch.PrivateMessage, client *twitch.Client) {
	var apiResponse string = utils.ApiNinjasRequest("facts")

	var facts []structs.Fact

	json.Unmarshal([]byte(apiResponse), &facts)

	if len(facts) > 0 {
		client.Say(message.Channel, facts[0].Fact+" forsenScoots")
	}
}
