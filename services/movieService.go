package services

import (
	"chillroom/configs"

	"github.com/ryanbradynd05/go-tmdb"
)

type MovieService interface {
	GetTrending() ([]interface{}, error)
	FindByID(movieID int) (*tmdb.Movie, error)
	SearchMovies(movieName string) (*tmdb.MovieSearchResults, error)
}

type movieService struct {
}

// getTrending implements MovieService
func (*movieService) GetTrending() ([]interface{}, error) {
	movie, err := configs.MovieConfig().GetTrending("movie", "day")
	var result []interface{}
	if err != nil {
		return result, err
	}
	for _, item := range movie.Results {
		data := map[string]interface{}{
			"id":            item.ID,
			"name":          item.Title,
			"summary":       item.Overview,
			"rating":        item.VoteAverage,
			"release_dates": item.ReleaseDate,
			"cover":         item.PosterPath,
		}
		result = append(result, data)
	}
	return result, nil
}

// FindByID implements MovieService
func (*movieService) FindByID(movieID int) (*tmdb.Movie, error) {
	var options = make(map[string]string)
	options["language"] = "id"
	movie, err := configs.MovieConfig().GetMovieInfo(movieID, options)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

// SearchMovies implements MovieService
func (*movieService) SearchMovies(movieName string) (*tmdb.MovieSearchResults, error) {
	var options = make(map[string]string)
	options["language"] = "id"
	movie, err := configs.MovieConfig().SearchMovie(movieName, options)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func NewMovieService() MovieService {
	return &movieService{}
}
