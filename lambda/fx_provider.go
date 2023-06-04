package lambda

import (
	"go.uber.org/fx"

	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/lambda/adapter"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/lambda/handler"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/lambda/runner"
)

var Module = fx.Options(
	adapter.Module,
	handler.Module,
	runner.Module,
)
