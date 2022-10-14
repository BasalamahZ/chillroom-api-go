package configs

import (
	"os"

	"github.com/Henry-Sarabia/igdb/v2"
	"github.com/ryanbradynd05/go-tmdb"
)

func MovieConfig() (tmdbAPI *tmdb.TMDb) {
	dsn := os.Getenv("MOVIE_API_KEY")
	config := tmdb.Config{
		APIKey:   dsn,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)
	return tmdbAPI
}

func GameConfig() (client *igdb.Client) {
	client = igdb.NewClient(os.Getenv("TWITCH_CLIENT_ID"), os.Getenv("ACCESS_TOKEN"), nil)
	return client
}
