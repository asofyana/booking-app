package services

import (
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
	ProcessLogin(c *gin.Context, username, password string) error
}

func (s *LoginService) ProcessLogin(c *gin.Context, username, password string) error {
	user, err := s.UserRepository.GetByUserName(username)
	if err != nil {
		return err
	}

	if user.UserName == "" {
		return fmt.Errorf("invalid username or password")
	}

	if user.Password != password {
		return fmt.Errorf("invalid username or password")
	}

	SetUserSession(c, user)

	return nil
}
