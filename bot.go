package main

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
)

func main() {
	fmt.Println("Init...")

	iniData := utils.GetIniData()

	channels := []string{
		iniData.Section("main").Key("main_channel").String(),
	}
	var movieList []structs.Movie

	client := twitch.NewClient("movisbot", "oauth:"+iniData.Section("main").Key("oauth_access_token").String())

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
		parseMovieSchedule(message, client, &movieList)
		parseCommand(message, client, movieList)
	})

	for _, channel := range channels {
		client.Join(channel)
	}

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
