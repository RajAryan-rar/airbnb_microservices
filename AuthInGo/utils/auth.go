package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	// "github.com/golang-jwt/jwt/v5"
)

func HashPassword(password string) (string, error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost);

	if err != nil {
		fmt.Println("Error hashing the password!");
		return "",err;
	}
	return string(hash), nil;
}

func CheckPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	return err == nil
}

