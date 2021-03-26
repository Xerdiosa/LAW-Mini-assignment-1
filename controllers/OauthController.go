package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Xerdiosa/LAW-Mini-assignment-1/helpers"
	"github.com/Xerdiosa/LAW-Mini-assignment-1/models"
	"github.com/Xerdiosa/LAW-Mini-assignment-1/services"
)

func InitOauthController(oauthService services.IOauthService) *OauthController {

	oauthController := new(OauthController)
	oauthController.OauthService = oauthService

	return oauthController
}

type OauthController struct {
	OauthService services.IOauthService
}

func (p *OauthController) Token(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	request := models.RequestToken{
		UserID:       req.FormValue("username"),
		UserPassword: req.FormValue("password"),
		ClientID:     req.FormValue("client_id"),
		ClientSecret: req.FormValue("client_secret"),
		GrantType:    req.FormValue("grant_type"),
	}

	acess, refresh, err := p.OauthService.Token(request)
	if err != nil {
		helpers.ResponseBadRequest(res, 400, err)
		return
	}
	response := helpers.APIResponseToken{
		AccessToken:  acess,
		ExpiresIn:    5,
		TokenType:    "Bearer",
		Scope:        nil,
		RefreshToken: refresh,
	}
	helpers.Response(res, response)
}

func (p *OauthController) Resource(res http.ResponseWriter, req *http.Request) {
	authHeader := req.Header.Get("Authorization")
	unparsedToken := strings.Split(authHeader, " ")
	if len(unparsedToken) != 2 {
		helpers.ResponseBadRequest(res, 401, errors.New("invalid authorization format"))
		return
	}
	parsedToken := unparsedToken[1]
	account, client, token, err := p.OauthService.Resource(parsedToken)
	if err != nil {
		helpers.ResponseBadRequest(res, 401, err)
		return
	}
	response := helpers.APIResponseResource{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ClientID:     client.ID,
		UserID:       account.UserId,
		FullName:     account.UserFullName,
		Npm:          account.NPM,
		Expires:      nil,
	}
	helpers.Response(res, response)
}
