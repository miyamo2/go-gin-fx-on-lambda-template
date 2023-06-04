package presentation

import (
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/presentation/controller"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/presentation/routing"
	"go.uber.org/fx"
)

var Module = fx.Options(controller.Module, routing.Module)
