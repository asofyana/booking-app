package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"
	"booking-app/internal/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(userId string) error
	UpdatePassword(c *gin.Context, oldPassword, newPassword, confirmPassword string) error
}

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAllUsers() ([]entity.User, error) {
	// TODO: Implement user listing logic
	return []entity.User{}, nil
}

func (s *UserService) CreateUser(user entity.User) error {
	// TODO: Implement user creation logic
	return nil
}

func (s *UserService) UpdateUser(user entity.User) error {
	// TODO: Implement user update logic
	return nil
}

func (s *UserService) DeleteUser(userId string) error {
	// TODO: Implement user deletion logic
	return nil
}

func (s *UserService) UpdatePassword(c *gin.Context, oldPassword, newPassword, confirmPassword string) error {

	minLength := 6
	if len(newPassword) < minLength {
		return fmt.Errorf("Password length minimum %d", minLength)
	}

	if newPassword != confirmPassword {
		return fmt.Errorf("new password should be same as confirm password")
	}

	user := GetUserSession(c)
	userDb, _ := s.userRepo.GetByUserName(user.UserName)

	if !utils.VerifyPassword(oldPassword, userDb.Password) {
		return fmt.Errorf("Invalid old password")
	}

	hashedPassword, _ := utils.HashPassword(newPassword)
	userDb.Password = hashedPassword
	err := s.userRepo.UpdatePassword(userDb)

	if err != nil {
		return fmt.Errorf("Error update password")
	}

	return nil
}
