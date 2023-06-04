package usecase

type SayByeUsecase interface {
	SayBye(username string) (string, error)
}
