package commands

import (
	"bot/internal/utils"
	"os"
	"path/filepath"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
)

// Prints a list of available bot commands
func Commands(message twitch.PrivateMessage, client *twitch.Client) {
	iniData := utils.GetIniData()

	files, err := os.ReadDir("./internal/commands")
	var commands []string

	if err == nil {
		for _, file := range files {
			var noExtensiionName string = (strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
			commands = append(commands, iniData.Section("main").Key("bot_command_prefix").String()+noExtensiionName)
		}
	}

	if len(commands) > 0 {
		client.Reply(message.Channel, message.ID, "Available commands: "+strings.Join(commands, ", "))
	}
}
