package routes

import (
	"net/http"

	"github.com/Xerdiosa/LAW-Mini-assignment-1/controllers"
	"github.com/Xerdiosa/LAW-Mini-assignment-1/services"

	"github.com/gorilla/mux"
)

// Route is
type Route struct{}

func (r *Route) Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	api := router.PathPrefix("/oauth").Subrouter()

	oauthService := new(services.OauthService)
	oauthController := controllers.InitOauthController(oauthService)

	api.HandleFunc("/token", oauthController.Token).Methods(http.MethodPost)
	api.HandleFunc("/resource", oauthController.Resource).Methods(http.MethodPost)

	return router
}
