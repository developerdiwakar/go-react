package controllers

import (
	"gofiber-api/services/tokens"
	"os"
	"time"
)

const (
	MsgInvalidRequestFormat = "Invalid Request Format"
	MsgInvalidUsername      = "Login Error: Invalid Username"
	MsgIncorrectPassword    = "Login Error: Password Incorrect."
)

func CreateJwtToken(userId int) (string, error) {
	tokenMaker, err := tokens.NewJWTMaker(os.Getenv("AUTH_KEY"))
	if err != nil {
		return "", err
	}
	token, err := tokenMaker.CreateToken(int(userId), time.Minute*30)
	if err != nil {
		return "", err
	}
	return token, nil
}
