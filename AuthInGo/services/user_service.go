package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
	env "AuthInGo/config/env"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() (string,error)
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

func (u *UserServiceImpl) LoginUser() (string,error) {
	email := "user1@gmail.com"
	password := "example_pass"

	//Step 1: Make a repo call to get the user by email
	user,err := u.userRepository.GetByEmail(email);

	if err != nil {
		fmt.Println("Error fetching the user!");
		return "", err
	}

	//Step-2: If user exists or not. If not exists, return error
	if(user == nil) {
		fmt.Println("No user found with given email")
		return "", fmt.Errorf("no user found with email :",email);
	}

	//Step-3: If user exists check the password using utils.CheckPasswordHash
	check := utils.CheckPasswordHash(password,user.Password)

	if(!check) {
		fmt.Println("Incorrect Password!");
		return "", nil;
	}

	//Step-4: If password matches, print the jwt token, else return saying incorrect password
	payload := jwt.MapClaims{
		"email" : user.Email,
		"username" : user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,payload);

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET","token")));
	
	if(err != nil) {
		fmt.Println("Error signing token :",err);
		return "", err;
	}
	fmt.Println("JWT Token :", tokenString);

	return tokenString, nil;
}
