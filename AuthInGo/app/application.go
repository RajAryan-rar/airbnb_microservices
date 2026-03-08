package app

import (
	"AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	dbConfig "AuthInGo/config/db"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

func NewConfig() *Config {
	port := config.GetString("PORT", ":8080")
	return &Config{
		Addr: port,
	}
} 

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

type Application struct {
	Config Config
}

func (app *Application) Run() error {
	db, err := dbConfig.SetupDB()
	
	if err != nil {
		fmt.Println("Error setting up database", err)
		return  err
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)
	

	server := &http.Server {
		Addr: app.Config.Addr,
		Handler: router.SetupRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	} 
	fmt.Println("Starting server on", app.Config.Addr)
	return server.ListenAndServe()
}