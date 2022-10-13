package services

import (
	"chillroom/models"
	"chillroom/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Create(user models.User) (models.User, error)
	FindByEmail(loginRequest models.LoginRequest) (models.User, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

// Create implements AuthService
func (as *authService) Create(user models.User) (models.User, error) {
	reqRegister := models.User{}
	reqRegister.Name = user.Name
	reqRegister.Username = user.Username
	reqRegister.Email = user.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return user, err
	}
	reqRegister.Password = string(hash)
	reqRegister.ConfirmPassword = user.ConfirmPassword

	newUser, err := as.authRepository.Create(reqRegister)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// FindByEmail implements AuthService
func (as *authService) FindByEmail(loginRequest models.LoginRequest) (models.User, error) {
	user, err := as.authRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewAuthService(authRepository *repositories.AuthRepository) AuthService {
	return &authService{
		authRepository: *authRepository,
	}
}
