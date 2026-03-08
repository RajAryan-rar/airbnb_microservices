package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error){
	cfg := mysql.NewConfig()
	cfg.User = env.GetString("DB_USER","root")
	cfg.Passwd = env.GetString("DB_PASSWORD","root");
	cfg.Net = env.GetString("DB_NET","tcp")
	cfg.Addr = env.GetString("DB_ADDR","")
	cfg.DBName = env.GetString("DBName","auth_dev");

	fmt.Println("Connecting to the database", cfg.DBName, cfg.FormatDSN())

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error while connecting to the db", err)
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error while pinging database", pingErr)
		return nil, pingErr
	}

	fmt.Println("Connected to the database successfully: ", cfg.DBName)

	return db, nil
}