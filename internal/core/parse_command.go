package core

import (
	"bot/internal/commands"
	"bot/internal/structs"
	"bot/internal/utils"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gempir/go-twitch-irc/v4"
)

// Parses the name of the command and executes its handler
func ParseCommand(
	message twitch.PrivateMessage,
	client *twitch.Client,
	movieList []structs.Movie,
	gptResponses map[string]*structs.GptResponseState,
	cooldowns map[string]*time.Time,
) {
	iniData := utils.GetIniData()

	var prefix string = iniData.Section("main").Key("bot_command_prefix").String()

	if strings.HasPrefix(message.Message, prefix) {
		_, i := utf8.DecodeRuneInString(message.Message)
		var noPrefixMessage string = strings.TrimSpace(message.Message[i:])
		var splitMessage []string = strings.Split(noPrefixMessage, " ")
		var command string = strings.ToLower(splitMessage[0])
		var args []string = splitMessage[1:]

		switch command {
		case "now":
			commands.Now(message, client, movieList)
		case "next":
			commands.Next(message, client, movieList, false)
		case "remaining":
			commands.Next(message, client, movieList, true)
		case "joke":
			commands.Joke(message, client)
		case "fact":
			commands.Fact(message, client, cooldowns)
		case "gpt":
			commands.Gpt(message, client, strings.Join(args, " "), gptResponses, cooldowns)
		case "continue":
			commands.Continue(message, client, gptResponses, cooldowns)
		case "joinchannel":
			if len(args) > 0 {
				commands.JoinChannel(message, client, args[0])
			}
		case "leavechannel":
			if len(args) > 0 {
				commands.LeaveChannel(message, client, args[0])
			}
		case "commands":
			commands.Commands(message, client)
		case "movie", "movi", "plot":
			commands.Plot(message, client, movieList, strings.Join(args, " "))
		case "rating":
			commands.Rating(message, client, movieList, strings.Join(args, " "))
		case "trivia":
			commands.Trivia(message, client, movieList, strings.Join(args, " "))
		case "raffle":
			var seconds string = "15"
			if len(args) > 0 {
				seconds = args[0]
			}
			commands.Raffle(message, client, seconds)
		case "pyramid":
			commands.Pyramid(message, client, args, cooldowns)
		}
	}
}
