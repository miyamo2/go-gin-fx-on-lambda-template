package usecase

import (
	"fmt"
)

type DoGreetUsecaseImpl struct{}

func (u *DoGreetUsecaseImpl) DoGreet(username string) (string, error) {
	fmt.Println("[DEBUG] START DoGreetUsecase#DoGreet")
	defer fmt.Println("[DEBUG] END DoGreetUsecase#DoGreet")
	return fmt.Sprintf("Nice to meet you, %s.", username), nil
}

func NewDoGreetUsecaseImpl() DoGreetUsecase {
	return &DoGreetUsecaseImpl{}
}
