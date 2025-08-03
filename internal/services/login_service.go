package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewLoginService(userRepository repository.UserRepositoryInterface) *LoginService {
	return &LoginService{
		UserRepository: userRepository,
	}
}

type LoginServiceInterface interface {
	ProcessLogin(c *gin.Context, username, password string) (entity.User, error)
}

func (s *LoginService) ProcessLogin(c *gin.Context, username, password string) (entity.User, error) {
	user, err := s.UserRepository.GetByUserName(username)
	if err != nil {
		return entity.User{}, err
	}

	if user.UserName == "" {
		return entity.User{}, fmt.Errorf("invalid username or password")
	}

	if user.Password != password {
		return entity.User{}, fmt.Errorf("invalid username or password")
	}

	SetUserSession(c, user)

	return user, nil
}
