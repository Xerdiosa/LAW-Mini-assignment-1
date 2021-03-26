package models

import "time"

type Client struct {
	ID     string
	Secret string
}

type Account struct {
	UserId       string
	UserFullName string
	NPM          string
	Password     string
}

type RequestToken struct {
	UserID       string
	UserPassword string
	ClientID     string
	ClientSecret string
	GrantType    string
}

type Token struct {
	UserID       string
	ValidUntil   time.Time
	AccessToken  string
	RefreshToken string
}

var ClientData = Client{
	ID:     "87647232352",
	Secret: "108358321",
}
var AccountData = Account{
	UserId:       "akbar",
	UserFullName: "Muhammad Aulia Akbar",
	NPM:          "1806133805",
	Password:     "password",
}

var TokenRepository map[string]Token = make(map[string]Token)
