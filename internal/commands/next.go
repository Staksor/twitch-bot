package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"fmt"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// Prints the name of the next movie and how much time left till it starts
func Next(
	message twitch.PrivateMessage,
	client *twitch.Client,
	movieList []structs.Movie,
	onlyRemainder bool,
) {
	if len(movieList) == 0 {
		client.Reply(message.Channel, message.ID, "eShrug")

		return
	}

	timezone, _ := time.LoadLocation("Europe/Berlin")
	time.Local = timezone

	var nextMovie structs.Movie
	var leftTime time.Duration
	var formatString string = ""
	var progressString string = ""
	var response string = ""
	var now = time.Now()
	var found bool = false

	for i, movie := range movieList {
		if now.Before(movie.Timestamp) && !found {
			nextMovie = movieList[i]
			leftTime = nextMovie.Timestamp.Sub(now)

			if onlyRemainder {
				formatString = "%s"
			} else {
				formatString = "(in %s)"
			}

			progressString = fmt.Sprintf(
				formatString,
				utils.FormatDuration(leftTime),
			)

			if onlyRemainder {
				response = progressString
			} else {
				response = nextMovie.Name + progressString
			}

			found = true
		}
	}

	if found {
		client.Reply(message.Channel, message.ID, response)
	} else {
		client.Reply(message.Channel, message.ID, "eShrug")
	}
}
