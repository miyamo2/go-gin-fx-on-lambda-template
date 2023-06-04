package usecase

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewSayHelloUsecaseImpl),
	fx.Provide(NewDoGreetUsecaseImpl),
	fx.Provide(NewSayByeUsecaseImpl),
)
