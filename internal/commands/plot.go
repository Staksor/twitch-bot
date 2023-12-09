package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/gocolly/colly"
)

// Prints the movie plot
func Plot(
	message twitch.PrivateMessage,
	client *twitch.Client,
	movieList []structs.Movie,
	movieName string,
) {
	var movieTitle string = ""
	var movieYear string = ""
	var moviePlot string = ""
	var movieDirector string = "???"
	var movieCast []string

	if len(movieName) == 0 {
		currentMovie, _ := utils.GetCurrentMovie(movieList)
		movieName = currentMovie.Name
	}
	movieName = strings.TrimSpace(movieName)

	scraper := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Go to the first result link on the search page
	scraper.OnHTML("[data-testid=\"find-results-section-title\"] > div > ul > li:first-child > div > div > a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		parsedUrl, _ := url.Parse(link)
		// Movie main page
		scraper.Visit(e.Request.AbsoluteURL(parsedUrl.EscapedPath()))
	})

	scraper.OnHTML("[data-testid=\"plot\"] span:first-child", func(e *colly.HTMLElement) {
		moviePlot = strings.Replace(e.Text, "Read all", "", -1)
	})
	scraper.OnHTML("h1[data-testid=\"hero__pageTitle\"] span", func(e *colly.HTMLElement) {
		movieTitle = e.Text
	})
	scraper.OnHTML("[data-testid=\"find-results-section-title\"] li:first-child a[href] + ul span", func(e *colly.HTMLElement) {
		movieYear = e.Text
	})
	scraper.OnHTML("[data-testid=\"title-pc-principal-credit\"]:first-child a", func(e *colly.HTMLElement) {
		movieDirector = " by " + e.Text
	})
	scraper.OnHTML("[data-testid=\"genres\"] ~ div [data-testid=\"title-pc-principal-credit\"]:last-child a.ipc-metadata-list-item__list-content-item", func(e *colly.HTMLElement) {
		movieCast = append(movieCast, e.Text)
	})

	scraper.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Host", "www.imdb.com")
		r.Headers.Set("Cookie", "lc-main=en_US")
		fmt.Println("Visiting", r.URL.String())
	})
	scraper.OnError(func(r *colly.Response, e error) {
		log.Println("Error:", e, r.Request.URL, string(r.Body))
	})

	scraper.Visit(fmt.Sprintf("https://www.imdb.com/find/?exact=true&q=%s", url.QueryEscape(movieName)))

	if len(moviePlot) > 0 {
		client.Reply(message.Channel, message.ID, fmt.Sprintf("%s (%s)%s. Starring %s. %s", movieTitle, movieYear, movieDirector, strings.Join(movieCast, ", "), moviePlot))
	} else {
		client.Reply(message.Channel, message.ID, "eShrug")
	}
}
