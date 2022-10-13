package configs

import (
	"os"

	"github.com/ryanbradynd05/go-tmdb"
)

func MovieConfig() (tmdbAPI *tmdb.TMDb){
	dsn := os.Getenv("MOVIE_API_KEY")
	config := tmdb.Config{
		APIKey:   dsn,
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)
	return tmdbAPI
}
