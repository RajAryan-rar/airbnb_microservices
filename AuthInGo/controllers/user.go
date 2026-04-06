package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)


type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserById called in UserController")
	uc.UserService.GetUserById()
	w.Write([]byte("user fetching endpoint"))
}

func (uc *UserController) CreateUserController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser called in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("user creation endpoint"))
}

func (uc * UserController) LoginUserController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginUser called in UserController")
	uc.UserService.LoginUser()
	w.Write([]byte("user login endpoint done"))
}