package services

import "testing"

func TestHello(t *testing.T) {
	homeService := new(HomeService)

	if homeService.HelloWorld() != "Hello World" {
		t.Errorf("Wanted Hello World, got %s", homeService.HelloWorld())
	}
}
