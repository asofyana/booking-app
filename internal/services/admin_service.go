package services

import (
	"booking-app/internal/entity"
	"booking-app/internal/repository"
)

type AdminServiceInterface interface {
	GetAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(userId string) error
}

type AdminService struct {
	userRepo repository.UserRepositoryInterface
}

func NewAdminService(userRepo repository.UserRepositoryInterface) *AdminService {
	return &AdminService{
		userRepo: userRepo,
	}
}

func (s *AdminService) GetAllUsers() ([]entity.User, error) {
	// TODO: Implement user listing logic
	return []entity.User{}, nil
}

func (s *AdminService) CreateUser(user entity.User) error {
	// TODO: Implement user creation logic
	return nil
}

func (s *AdminService) UpdateUser(user entity.User) error {
	// TODO: Implement user update logic
	return nil
}

func (s *AdminService) DeleteUser(userId string) error {
	// TODO: Implement user deletion logic
	return nil
}
