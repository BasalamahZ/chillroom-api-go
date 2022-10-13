package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string `binding:"required"`
	Username        string `gorm:"unique" binding:"required"`
	Email           string `gorm:"unique" binding:"required,email"`
	Password        string `binding:"required,min=6,eqfield=ConfirmPassword"`
	ConfirmPassword string `gorm:"-" binding:"required,min=6,eqfield=Password"`
}

type LoginRequest struct {
	Email           string `binding:"required,email"`
	Password        string `binding:"required,min=6"`
}
