package repositories

import (
	"chillroom/configs"
	"chillroom/models"
)

type AuthRepository interface {
	Create(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type authRepository struct {
}

// Create implements AuthRepository
func (*authRepository) Create(user models.User) (models.User, error) {
	db := configs.DB.Create(&user)
	if db != nil {
		return user, db.Error
	}
	return user, nil
}

// FindByEmail implements AuthRepository
func (*authRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	db := configs.DB.First(&user, "email = ?", email)
	if db != nil {
		return user, db.Error
	}
	return user, nil
}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}
