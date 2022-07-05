package main

import "fmt"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//go:generate moq -out main_mock_test.go . UserRepositoryInterface
type UserRepositoryInterface interface {
	GetAllUsers() ([]User, error)
}

type UserService struct {
	UserRepositoryInterface
}

func (s UserService) GetUser() ([]User, error) {
	users, err := s.UserRepositoryInterface.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for i := range users {
		users[i].Password = "*****"
	}
	return users, nil
}

type UserRepository struct{}

func (r UserRepository) GetAllUsers() ([]User, error) {
	users := []User{
		{"real", "real"},
		{"real2", "real2"},
	}
	return users, nil
}

func main() {
	repository := UserRepository{}
	service := UserService{repository}
	users, _ := service.GetUser()
	fmt.Println(users)
}
