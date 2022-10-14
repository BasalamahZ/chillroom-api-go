package controllers

import (
	"chillroom/cache"
	"chillroom/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	movieService services.MovieService
	apiCache   cache.ApiCache
}

func NewMovieController(movieService *services.MovieService, cache *cache.ApiCache) MovieController {
	return MovieController{
		movieService: *movieService,
		apiCache: *cache,
	}
}

func (mc *MovieController) GetTrending(c *gin.Context) {
	var movies []interface{} = mc.apiCache.Get("movies_trending")
	if movies == nil {
		movies, err := mc.movieService.GetTrending()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
			})
			return
		}
		mc.apiCache.Set("movies_trending", movies)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    movies,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    movies,
		})
	}
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
	movie, err := mc.movieService.SearchMovies(movieName)
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
