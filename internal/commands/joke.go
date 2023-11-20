package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"encoding/json"

	"github.com/gempir/go-twitch-irc/v4"
)

// Prints a random dad joke from API Ninjas
func Joke(message twitch.PrivateMessage, client *twitch.Client) {
	var apiResponse string = utils.ApiNinjasRequest("dadjokes")

	var jokes []structs.Joke

	json.Unmarshal([]byte(apiResponse), &jokes)

	if len(jokes) > 0 {
		client.Say(message.Channel, jokes[0].Joke+" Pepepains")
	}
}
