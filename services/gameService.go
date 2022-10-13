package services

import (
	"os"
	"time"

	"github.com/Henry-Sarabia/igdb/v2"
)

type GameService interface {
	GetTrending() ([]interface{}, error)
}

type gameService struct {
}

// GetTrending implements GameService
func (*gameService) GetTrending() ([]interface{}, error) {
	client := igdb.NewClient(os.Getenv("TWITCH_CLIENT_ID"), os.Getenv("ACCESS_TOKEN"), nil)
	byPop := igdb.ComposeOptions(
		igdb.SetLimit(50),
		igdb.SetFields("*"),
		igdb.SetOrder("hypes", igdb.OrderDescending),
		igdb.SetFilter("first_release_date", igdb.OpGreaterThan, "1658188700"),
		igdb.SetFilter("rating", igdb.OpGreaterThan, "75"),
	)
	games, err := client.Games.Index(byPop)
	var result []interface{}
	if err != nil {
		return result, err
	}
	for _, item := range games {
		cover, err := client.Covers.Get(item.Cover, igdb.SetFields("image_id"))
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

func NewGameService() GameService {
	return &gameService{}
}
