package controllers

import (
	"chillroom/helpers"
	"chillroom/models"
	"chillroom/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService *services.AuthService) AuthController {
	return AuthController{
		authService: *authService,
	}
}

func (ac *AuthController) Register(c *gin.Context){
	reqRegister := new(models.User)
	err := c.BindJSON(&reqRegister)
	if err != nil {
		errors := helpers.FormatError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors,
		})
		return
	}
	user, err := ac.authService.Create(*reqRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}
	// (-) hiding confirm password in response
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data" : user,
	})
}	

func (ac *AuthController) Login(c *gin.Context){
	reqLogin := new(models.LoginRequest)
	err := c.BindJSON(&reqLogin)
	if err != nil {
		errors := helpers.FormatError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors,
		})
		return
	}
	user, err := ac.authService.FindByEmail(*reqLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}
	// (-) hiding confirm password in response
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data" : user,
	})

}