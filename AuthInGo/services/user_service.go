package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
	"AuthInGo/utils"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("fetching user in UserService")
	u.userRepository.DeleteById(2)
	return nil
}

func (u *UserServiceImpl) CreateUser() error {

	password := "example_pass"
	
	hashedPassword, err := utils.HashPassword(password);

	if err != nil {
		fmt.Println("Error getting hashed password");
	}

	u.userRepository.Create(
		"user1_example",
		"user1@gmail.com",
		hashedPassword,
	)

	return  nil;
}	

func (u *UserServiceImpl) LoginUser() error {
	response := utils.CheckPasswordHash("example_pass_wrong","$2a$10$1sjQbYiJy4HbK4CEVRP40OCvdebx3OlBLUMW5CKFaZ143hD7lwPiy")
	fmt.Println("Login response :",response)
	return nil;
}
