package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"
	"booking-app/internal/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	DeleteUser(userId string) error
	UpdatePassword(c *gin.Context, oldPassword, newPassword, confirmPassword string) error
	SearchUsers(user entity.User) ([]entity.User, error)
	GetUserById(userid int) (entity.User, error)
	ResetPassword(userid int) error
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

func (s *UserService) DeleteUser(userId string) error {
	// TODO: Implement user deletion logic
	return nil
}

func (s *UserService) CreateUser(c *gin.Context) error {
	user := GetUserSession(c)

	fullName := c.PostForm("fullName")
	username := c.PostForm("username")
	title := c.PostForm("title")
	email := c.PostForm("email")
	phoneNumber := c.PostForm("phoneNumber")
	role := c.PostForm("role")
	status := c.PostForm("status")

	newUser := entity.User{
		FullName:    fullName,
		UserName:    username,
		Title:       title,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        role,
		Status:      status,
		Password:    utils.GetConfig().DefaultPassword,
		CreatedBy:   user.UserName,
	}

	return s.userRepo.CreateUser(newUser)
}

func (s *UserService) UpdateUser(c *gin.Context) error {
	user := GetUserSession(c)

	fullName := c.PostForm("fullName")
	username := c.PostForm("username")
	title := c.PostForm("title")
	email := c.PostForm("email")
	phoneNumber := c.PostForm("phoneNumber")
	role := c.PostForm("role")
	status := c.PostForm("status")
	userId, _ := strconv.Atoi(c.Request.FormValue("userid"))

	newUser := entity.User{
		UserId:      userId,
		FullName:    fullName,
		UserName:    username,
		Title:       title,
		Email:       email,
		PhoneNumber: phoneNumber,
		Role:        role,
		Status:      status,
		UpdatedBy:   user.UserName,
	}

	return s.userRepo.UpdateUser(newUser)
}

func (s *UserService) UpdatePassword(c *gin.Context, oldPassword, newPassword, confirmPassword string) error {

	minLength := 6
	if len(newPassword) < minLength {
		return fmt.Errorf("password length minimum %d", minLength)
	}

	if newPassword != confirmPassword {
		return fmt.Errorf("new password should be same as confirm password")
	}

	user := GetUserSession(c)
	userDb, _ := s.userRepo.GetByUserName(user.UserName)

	if !utils.VerifyPassword(oldPassword, userDb.Password) {
		return fmt.Errorf("invalid old password")
	}

	hashedPassword, _ := utils.HashPassword(newPassword)
	userDb.Password = hashedPassword
	err := s.userRepo.UpdatePassword(userDb)

	if err != nil {
		return fmt.Errorf("error update password")
	}

	return nil
}

func (s *UserService) SearchUsers(user entity.User) ([]entity.User, error) {
	return s.userRepo.SearchUsers(user)
}

func (s *UserService) GetUserById(userid int) (entity.User, error) {
	return s.userRepo.GetByUserId(userid)
}

func (s *UserService) ResetPassword(userid int) error {
	var user entity.User
	user.UserId = userid
	user.Password = utils.GetConfig().DefaultPassword
	return s.userRepo.UpdatePassword(user)
}
