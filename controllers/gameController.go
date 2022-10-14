package controllers

import (
	"chillroom/cache"
	"chillroom/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	gameService services.GameService
	apiCache    cache.ApiCache
}

func NewGameController(gameService *services.GameService, cache *cache.ApiCache) GameController {
	return GameController{
		gameService: *gameService,
		apiCache:    *cache,
	}
}

func (gc *GameController) GetTrending(c *gin.Context) {
	var games []interface{} = gc.apiCache.Get("games_trending")
	if games == nil {
		games, err := gc.gameService.GetTrending()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			return
		}
		gc.apiCache.Set("games_trending", games)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    games,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    games,
		})
	}
}

func (gc *GameController) FindByID(c *gin.Context) {
	gameID := c.Param("id")
	newGameID, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to convert id to integer",
		})
		return
	}
	games, err := gc.gameService.FindByID(newGameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    games,
	})
}

func (gc *GameController) SearchGames(c *gin.Context) {
	gameName := c.Query("title")
	games, err := gc.gameService.SearchGames(gameName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    games,
	})
}