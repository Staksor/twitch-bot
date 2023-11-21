package main

import (
	"bot/internal/core"
	"bot/internal/structs"
	"bot/internal/utils"
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

func main() {
	fmt.Println("Starting...")

	iniData := utils.GetIniData()

	channels := []string{
		iniData.Section("main").Key("main_channel").String(),
		iniData.Section("main").Key("bot_account_name").String(),
	}
	var movieList []structs.Movie
	cooldowns := make(map[string]*time.Time)
	gptResponses := make(map[string]*structs.GptResponseState)

	client := twitch.NewClient(iniData.Section("main").Key("bot_account_name").String(), "oauth:"+iniData.Section("main").Key("oauth_access_token").String())

	client.OnConnect(func() {
		fmt.Println("Connected...")
	})

	client.OnSelfJoinMessage(func(message twitch.UserJoinMessage) {
		fmt.Printf("Joined %s\n", message.Channel)
		client.Say(message.Channel, "!time")
	})

	client.OnSelfPartMessage(func(message twitch.UserPartMessage) {
		fmt.Printf("Left %s\n", message.Channel)
	})

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		core.ParseMovieSchedule(message, client, &movieList)
		core.ParseCommand(message, client, movieList, gptResponses, cooldowns)
	})

	for _, channel := range channels {
		client.Join(channel)
	}

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
