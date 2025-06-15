package repository

import (
	"booking-app/internal/entity"
	"database/sql"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

type UserRepositoryInterface interface {
	GetByUserName(username string) (entity.User, error)
}

func (s *UserRepository) GetByUserName(username string) (entity.User, error) {
	result := s.DB.QueryRow("select user_id, username, password, full_name from users where username = ?", username)

	var user entity.User

	err := result.Scan(&user.UserId, &user.UserName, &user.Password, &user.FullName)
	if err != nil {
		log.Println("Error scanning user: ", err.Error())
		return entity.User{}, err
	}

	return user, nil
}
