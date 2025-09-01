package repository

import (
	"booking-app/internal/entity"
	"database/sql"
	"fmt"
	"log"
	"strings"
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
	GetAllUsers() ([]entity.User, error)
	CreateUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(userId string) error
	UpdatePassword(user entity.User) error
	SearchUsers(user entity.User) ([]entity.User, error)
}

func (s *UserRepository) GetByUserName(username string) (entity.User, error) {
	result := s.DB.QueryRow("select user_id, username, password, full_name, role from users where username = ?", username)

	var user entity.User

	err := result.Scan(&user.UserId, &user.UserName, &user.Password, &user.FullName, &user.Role)
	if err != nil {
		log.Println("Error scanning user: ", err.Error())
		return entity.User{}, err
	}

	return user, nil
}

func (s *UserRepository) GetAllUsers() ([]entity.User, error) {
	rows, err := s.DB.Query("select user_id, username, password, full_name, email, phone_number, role from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.UserId, &user.UserName, &user.Password, &user.FullName, &user.Email, &user.PhoneNumber, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserRepository) CreateUser(user entity.User) error {
	_, err := s.DB.Exec("insert into users (username, password, full_name, email, phone_number, role) values (?, ?, ?, ?, ?, ?)",
		user.UserName, user.Password, user.FullName, user.Email, user.PhoneNumber, user.Role)
	return err
}

func (s *UserRepository) UpdateUser(user entity.User) error {
	_, err := s.DB.Exec("update users set username=?, password=?, full_name=?, email=?, phone_number=?, role=? where user_id=?",
		user.UserName, user.Password, user.FullName, user.Email, user.PhoneNumber, user.Role, user.UserId)
	return err
}

func (s *UserRepository) UpdatePassword(user entity.User) error {
	_, err := s.DB.Exec("update users set password=? where user_id=?", user.Password, user.UserId)
	return err
}

func (s *UserRepository) DeleteUser(userId string) error {
	_, err := s.DB.Exec("delete from users where user_id = ?", userId)
	return err
}

func (s *UserRepository) SearchUsers(user entity.User) ([]entity.User, error) {

	fmt.Println("========== Search User")

	sql := "select user_id, username, password, full_name, IFNULL(email,'') as email, IFNULL(phone_number,'') as phone_number , role from users"
	conditions := []string{}
	params := []interface{}{}

	if user.FullName != "" {
		conditions = append(conditions, "lower(full_name) like '%?%'")
		params = append(params, strings.ToLower(user.FullName))
	}

	if user.UserName != "" {
		conditions = append(conditions, "lower(username) like '%?%'")
		params = append(params, strings.ToLower(user.UserName))
	}

	if len(conditions) > 0 {
		sql += " WHERE " + strings.Join(conditions, " AND ")
	}

	fmt.Println("sql:", sql)
	fmt.Println(params)

	rows, err := s.DB.Query(sql, params...)

	if err != nil {
		fmt.Println("query db failed")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	fmt.Println("query db success")

	var users []entity.User
	for rows.Next() {
		fmt.Println("NEXT")
		var user entity.User
		err := rows.Scan(&user.UserId, &user.UserName, &user.Password, &user.FullName, &user.Email, &user.PhoneNumber, &user.Role)
		if err != nil {
			fmt.Println("scan failed")
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
