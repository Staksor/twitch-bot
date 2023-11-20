package main

import (
	"bot/internal/commands"
	"bot/internal/structs"
	"strings"
	"unicode/utf8"

	"github.com/gempir/go-twitch-irc/v4"
)

func parseCommand(message twitch.PrivateMessage, client *twitch.Client, movieList []structs.Movie) {
	const prefix string = "!"

	if strings.HasPrefix(message.Message, prefix) {
		_, i := utf8.DecodeRuneInString(message.Message)
		var noPrefixMessage string = message.Message[i:]
		var splitMessage []string = strings.Split(noPrefixMessage, " ")
		var command string = splitMessage[0]
		var args []string = splitMessage[1:]

		switch {
		case command == "now":
			commands.Now(message, client, movieList)
		case command == "next":
			commands.Next(message, client, movieList)
		case command == "test":
			commands.Test(message, client)
		case command == "joke":
			commands.Joke(message, client)
		case command == "fact":
			commands.Fact(message, client)
		case command == "gpt":
			commands.Gpt(message, client, strings.Join(args, " "))
		case command == "joinchannel":
			commands.JoinChannel(message, client, args[0])
		case command == "leavechannel":
			commands.LeaveChannel(message, client, args[0])
		}
	}
}
