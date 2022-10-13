package controllers

import (
	"chillroom/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieService services.MovieService
}

func NewMovieController(movieService *services.MovieService) MovieController {
	return MovieController{
		movieService: *movieService,
	}
}

func (mc *MovieController) GetTrending(c *gin.Context) {
	movie, err := mc.movieService.GetTrending()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    movie,
	})
}

func (mc *MovieController) FindByID(c *gin.Context) {
	movieID := c.Param("id")
	newMovieID, err := strconv.Atoi(movieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to convert id to integer",
		})
		return
	}
	movie, err := mc.movieService.FindByID(newMovieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    movie,
	})
}

func (mc *MovieController) SearchMovie(c *gin.Context) {
	movieName := c.Query("query")
	movie, err := mc.movieService.SearchMovie(movieName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    movie,
	})
}