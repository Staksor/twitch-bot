package commands

import (
	"bot/internal/structs"
	"bot/internal/utils"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strings"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/gocolly/colly"
	strip "github.com/grokify/html-strip-tags-go"
)

// Prints the movie plot
func Trivia(
	message twitch.PrivateMessage,
	client *twitch.Client,
	movieList []structs.Movie,
	movieName string,
) {
	var movieTrivia []string

	if len(movieName) == 0 {
		currentMovie, _ := utils.GetCurrentMovie(movieList)
		movieName = currentMovie.Name
	}
	movieName = strings.TrimSpace(movieName)

	scraper := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	// Go to the first result link on the search page
	scraper.OnHTML("[data-testid=\"find-results-section-title\"] > div > ul > li:first-child a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		parsedUrl, _ := url.Parse(link)
		// Movie main page
		scraper.Visit(e.Request.AbsoluteURL(parsedUrl.EscapedPath()) + "trivia")
	})

	scraper.OnHTML("section.ipc-page-section.ipc-page-section--base [data-testid=\"item-id\"] .ipc-html-content-inner-div", func(e *colly.HTMLElement) {
		var currentTrivia string = strip.StripTags(e.Text)

		if len(currentTrivia) <= 460 {
			movieTrivia = append(movieTrivia, currentTrivia)
		}
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

	if len(movieTrivia) > 0 {
		client.Reply(message.Channel, message.ID, movieTrivia[rand.Intn(len(movieTrivia))])
	} else {
		client.Reply(message.Channel, message.ID, "eShrug")
	}
}
