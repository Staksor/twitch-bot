package core

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

func ParseMovieSchedule(message twitch.PrivateMessage, client *twitch.Client, MovieList *[]structs.Movie) {
	iniData := utils.GetIniData()

	botName := iniData.Section("main").Key("schedule_bot_name").String()

	if message.User.Name == botName {
		r, _ := regexp.Compile("\\[.+?UTC\\+\\d\\]")

		if !r.MatchString(message.Message) {
			return
		}

		var movieListString string = r.ReplaceAllString(message.Message, "")

		timezone, err := time.LoadLocation("Europe/Berlin")

		if err != nil {
			fmt.Println("Error loading location:", err)
			return
		}

		time.Local = timezone

		var now time.Time = time.Now()
		var nowString string = now.Format("2006-01-02 ")

		if len(movieListString) > 1 {
			*MovieList = nil
			var parsedMovieStrings []string = strings.Split(movieListString, "⏩")

			for _, movieString := range parsedMovieStrings {
				r, _ := regexp.Compile("(.+?)\\((\\d{1,2}:\\d{2})\\)")
				var parsedMovieString []string = r.FindStringSubmatch(movieString)

				movieName := parsedMovieString[1]
				movieTimestamp, _ := time.ParseInLocation("2006-01-02 15:04", nowString+parsedMovieString[2], timezone)

				parsedMovie := structs.Movie{
					Name:      movieName,
					Timestamp: movieTimestamp,
				}
				*MovieList = append(*MovieList, parsedMovie)
			}
		}
	}
}
