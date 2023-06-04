package application

import (
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/application/usecase"
	"go.uber.org/fx"
)

var Module = fx.Options(usecase.Module)
