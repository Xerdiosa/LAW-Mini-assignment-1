package services

type IHomeService interface {
	HelloWorld() string
}

type HomeService struct {
}

func (s HomeService) HelloWorld() string {
	return "Hello World"
}
