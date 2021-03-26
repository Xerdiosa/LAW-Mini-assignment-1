package services

import (
	"errors"
	"time"

	"github.com/Xerdiosa/LAW-Mini-assignment-1/helpers"
	"github.com/Xerdiosa/LAW-Mini-assignment-1/models"
)

type IOauthService interface {
	Token(models.RequestToken) (string, string, error)
	Resource(string) (models.Account, models.Client, models.Token, error)
}

type OauthService struct {
}

func (s OauthService) Token(request models.RequestToken) (access string, refresh string, err error) {
	if request.GrantType != "password" {
		err := errors.New("invalid grant type")
		return access, refresh, err
	}
	if models.ClientData.ID != request.ClientID || models.ClientData.Secret != request.ClientSecret {
		err := errors.New("client secret is wrong or client does not exist")
		return access, refresh, err
	}
	if models.AccountData.UserId != request.UserID || models.AccountData.Password != request.UserPassword {
		err := errors.New("password is wrong or user does not exist")
		return access, refresh, err
	}

	access = helpers.GenerateToken()
	refresh = helpers.GenerateToken()

	models.TokenRepository[access] = models.Token{
		UserID:       request.ClientID,
		ValidUntil:   time.Now().Add(time.Minute * 5),
		AccessToken:  access,
		RefreshToken: refresh,
	}

	return access, refresh, err
}

func (s OauthService) Resource(access string) (account models.Account, client models.Client, token models.Token, err error) {
	token, ok := models.TokenRepository[access]
	if !ok {
		err = errors.New("invalid or expired token")
		return account, client, token, err
	}
	if token.ValidUntil.Before(time.Now()) {
		err = errors.New("invalid or expired token")
		return account, client, token, err
	}
	return models.AccountData, models.ClientData, token, err
}
