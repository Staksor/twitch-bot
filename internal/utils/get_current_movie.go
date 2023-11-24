package utils

import (
	"bot/internal/structs"
	"time"
)

func GetCurrentMovie(movieList []structs.Movie) (structs.Movie, int) {
	var currentMovie structs.Movie
	var currentMovieIndex int
	timezone, _ := time.LoadLocation("Europe/Berlin")
	time.Local = timezone
	var now = time.Now()

	for i, movie := range movieList {
		if now.After(movie.Timestamp) {
			currentMovie = movie
			currentMovieIndex = i
		}
	}

	return currentMovie, currentMovieIndex
}
