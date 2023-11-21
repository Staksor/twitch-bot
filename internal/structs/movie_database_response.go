package structs

type MovieDatabaseTitle struct {
	Text string `json:"text"`
}

type MovieDatabasePlot struct {
	PlotText MovieDatabasePlotText `json:"plotText"`
}

type MovieDatabasePlotText struct {
	PlainText string `json:"plainText"`
}

type MovieDatabaseReleaseYear struct {
	Year int `json:"year"`
}

type MovieDatabaseData struct {
	TitleText   MovieDatabaseTitle       `json:"titleText"`
	Plot        MovieDatabasePlot        `json:"plot"`
	ReleaseYear MovieDatabaseReleaseYear `json:"releaseYear"`
}

type MovieDatabaseResponse struct {
	Results []MovieDatabaseData `json:"results"`
}
