package services

import (
	"chillroom/configs"
	"time"

	"github.com/Henry-Sarabia/igdb/v2"
)

type GameService interface {
	GetTrending() ([]interface{}, error)
	FindByID(gameID int) ([]interface{}, error)
	SearchGames(gameName string) ([]interface{}, error)
}

type gameService struct {
}

// GetTrending implements GameService
func (*gameService) GetTrending() ([]interface{}, error) {
	opts := igdb.ComposeOptions(
		igdb.SetLimit(50),
		igdb.SetFields("*"),
		igdb.SetOrder("hypes", igdb.OrderDescending),
		igdb.SetFilter("first_release_date", igdb.OpGreaterThan, "1658188700"),
		igdb.SetFilter("rating", igdb.OpGreaterThan, "75"),
	)
	games, err := configs.GameConfig().Games.Index(opts)
	var result []interface{}
	if err != nil {
		return result, err
	}
	for _, item := range games {
		cover, err := configs.GameConfig().Covers.Get(item.Cover, igdb.SetFields("image_id"))
		if err != nil {
			return result, err
		}
		img, err := cover.SizedURL(igdb.Size1080p, 1) // resize to largest image available
		if err != nil {
			return result, err
		}
		layout := "Jan 2, 2006 at 3:04pm (MST)"
		timeT := time.Unix(int64(item.FirstReleaseDate), 0).Format(layout)
		data := map[string]interface{}{
			"id":            item.ID,
			"name":          item.Name,
			"summary":       item.Summary,
			"rating":        item.Rating,
			"release_dates": timeT,
			"cover":         img,
		}
		result = append(result, data)
	}

	return result, nil
}

// FindByID implements GameService
func (*gameService) FindByID(gameID int) ([]interface{}, error) {
	games, err := configs.GameConfig().Games.Get(
		gameID,
		igdb.SetFields("*"),
	)
	var result []interface{}
	if err != nil {
		return result, err
	}
	result = append(result, games)
	return result, nil
}

// SearchGames implements GameService
func (*gameService) SearchGames(gameName string) ([]interface{}, error) {
	games, err := configs.GameConfig().Games.Search(
		gameName,
		igdb.SetFields("*"),
	)
	var result []interface{}
	if err != nil {
		return result, err
	}
	result = append(result, games)
	return result, nil
}

func NewGameService() GameService {
	return &gameService{}
}
