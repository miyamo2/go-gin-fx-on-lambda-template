package fxapp

import (
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/application"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/presentation"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/fxapp/runner"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/lambda"
	"go.uber.org/fx"
)

var FxApp *fx.App

func init() {
	FxApp = fx.New(
		fx.Options(
			application.Module,
			presentation.Module,
			lambda.Module,
		),
		fx.Invoke(
			func(lifecycle fx.Lifecycle, runner runner.FxAppRunner) {
				runner.Run(lifecycle)
			},
		),
	)
}
