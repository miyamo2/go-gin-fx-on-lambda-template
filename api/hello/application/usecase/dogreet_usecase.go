package usecase

type DoGreetUsecase interface {
	DoGreet(username string) (string, error)
}
