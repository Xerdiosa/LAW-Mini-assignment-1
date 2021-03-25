package controllers

import (
	"net/http"

	"github.com/Xerdiosa/LAW-Mini-assignment-1/helpers"
	"github.com/Xerdiosa/LAW-Mini-assignment-1/services"
)

func InitHomeController(homeService services.IHomeService) *HomeController {

	homeController := new(HomeController)
	homeController.HomeService = homeService

	return homeController
}

type HomeController struct {
	HomeService services.IHomeService
}

func (p *HomeController) HelloWorld(res http.ResponseWriter, req *http.Request) {

	hello := p.HomeService.HelloWorld()
	helpers.Response(res, http.StatusOK, hello)
}
