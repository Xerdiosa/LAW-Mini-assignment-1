package routes

import (
	"github.com/Xerdiosa/LAW-Mini-assignment-1/controllers"
	"github.com/Xerdiosa/LAW-Mini-assignment-1/services"

	"github.com/gorilla/mux"
)

// Route is
type Route struct{}

func (r *Route) Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	api := router.PathPrefix("/api").Subrouter()

	homeService := new(services.HomeService)
	homeController := controllers.InitHomeController(homeService)

	api.HandleFunc("/", homeController.HelloWorld).Methods("GET")

	return router
}
