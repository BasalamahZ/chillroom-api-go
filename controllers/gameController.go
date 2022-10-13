package controllers

import (
	"chillroom/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	gameService services.GameService
}

func NewGameController(gameService *services.GameService) GameController {
	return GameController{
		gameService: *gameService,
	}
}

func (gc *GameController) GetTrending(c *gin.Context) {
	game, err := gc.gameService.GetTrending()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    game,
	})
}
