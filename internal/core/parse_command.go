package core

import (
	"bot/internal/commands"
	"bot/internal/structs"
	"bot/internal/utils"
	"strings"
	"unicode/utf8"

	"github.com/gempir/go-twitch-irc/v4"
)

// Parses the name of the command and executes its handler
func ParseCommand(message twitch.PrivateMessage, client *twitch.Client, movieList []structs.Movie) {
	iniData := utils.GetIniData()

	var prefix string = iniData.Section("main").Key("bot_command_prefix").String()

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
		case command == "commands":
			commands.Commands(message, client)
		}
	}
}
