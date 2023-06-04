package usecase

import (
	"fmt"
)

type SayByeUsecaseImpl struct{}

func (u *SayByeUsecaseImpl) SayBye(username string) (string, error) {
	fmt.Println("[DEBUG] START SayByeUsecase#SayBye")
	defer fmt.Println("[DEBUG] END SayByeUsecase#SayBye")
	return fmt.Sprintf("Bye %s.", username), nil
}

func NewSayByeUsecaseImpl() SayByeUsecase {
	return &SayByeUsecaseImpl{}
}
