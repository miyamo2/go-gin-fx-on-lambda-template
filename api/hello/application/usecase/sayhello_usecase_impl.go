package usecase

import (
	"fmt"
)

type SayHelloUsecaseImpl struct{}

func (u *SayHelloUsecaseImpl) SayHello() (string, error) {
	fmt.Println("[DEBUG] START SayHelloUsecase#SayHello")
	defer fmt.Println("[DEBUG] END SayHelloUsecase#SayHello")
	return "Hello", nil
}

func NewSayHelloUsecaseImpl() SayHelloUsecase {
	return &SayHelloUsecaseImpl{}
}
