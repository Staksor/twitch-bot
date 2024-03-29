package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"fmt"
	"math"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
)

// Prints the name of the current movie, its current progress and total movie length
func Now(message twitch.PrivateMessage, client *twitch.Client, movieList []structs.Movie) {
	if len(movieList) == 0 {
		client.Reply(message.Channel, message.ID, "eShrug")

		return
	}

	timezone, _ := time.LoadLocation("Europe/Berlin")
	time.Local = timezone

	var currentMovie structs.Movie
	var nextMovie structs.Movie
	var progressTime time.Duration
	var totalTime time.Duration
	var progressString string = ""
	var now = time.Now()

	currentMovie, currentMovieIndex := utils.GetCurrentMovie(movieList)

	if len(movieList) > currentMovieIndex+1 {
		nextMovie = movieList[currentMovieIndex+1]
		progressTime = now.Sub(currentMovie.Timestamp)
		totalTime = nextMovie.Timestamp.Sub(currentMovie.Timestamp)

		progressString = fmt.Sprintf(
			"(%s/%s, %d%%)",
			utils.FormatDuration(progressTime),
			utils.FormatDuration(totalTime),
			int(math.Ceil(progressTime.Seconds()/totalTime.Seconds()*100)),
		)
	} else {
		progressString = ""
	}

	client.Reply(message.Channel, message.ID, currentMovie.Name+progressString)
}
