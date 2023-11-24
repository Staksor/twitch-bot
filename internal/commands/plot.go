package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Prints the movie plot
func Plot(
	message twitch.PrivateMessage,
	client *twitch.Client,
	movieList []structs.Movie,
	movieName string,
) {
	if len(movieName) == 0 {
		currentMovie, _ := utils.GetCurrentMovie(movieList)
		movieName = currentMovie.Name
	}
	movieName = strings.TrimSpace(movieName)

	iniData := utils.GetIniData()

	httpClient := &http.Client{}
	caser := cases.Title(language.English)
	req, _ := http.NewRequest("GET", "https://moviesdatabase.p.rapidapi.com/titles/search/title/"+caser.String(movieName)+"?exact=true&info=base_info", nil)
	req.Header.Set("X-RapidAPI-Key", iniData.Section("api").Key("movies_db_key").String())
	req.Header.Set("X-RapidAPI-Host", "moviesdatabase.p.rapidapi.com")
	res, _ := httpClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	bodyString := string(body)

	var results structs.MovieDatabaseResponse

	if err := json.Unmarshal([]byte(bodyString), &results); err == nil {
		if len(results.Results) > 0 {
			for i, movie := range results.Results {
				movieTitle := movie.TitleText.Text
				movieYear := movie.ReleaseYear.Year
				moviePlot := movie.Plot.PlotText.PlainText

				if len(moviePlot) > 0 {
					if movieYear > 0 {
						client.Reply(message.Channel, message.ID, fmt.Sprintf("%s (%d). %s", movieTitle, movieYear, moviePlot))
					} else {
						client.Reply(message.Channel, message.ID, fmt.Sprintf("%s. %s", movieTitle, moviePlot))
					}

					break
				} else if (len(results.Results) - 1) == i {
					client.Reply(message.Channel, message.ID, "eShrug")
				}
			}
		} else {
			client.Reply(message.Channel, message.ID, "eShrug")
		}
	} else {
		fmt.Println(err)
	}
}
