package controllers

import (
	"gofiber-api/services/tokens"
	"os"
	"strconv"
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
	expiryTime, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRY_TIME"))
	token, err := tokenMaker.CreateToken(int(userId), time.Minute*time.Duration(expiryTime))
	if err != nil {
		return "", err
	}
	return token, nil
}
