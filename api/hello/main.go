package main

import (
	"fmt"

	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/configure/fxapp"
	"go.uber.org/fx"
)

var app *fx.App

func init() {
	app = fxapp.FxApp
}

func main() {
	fmt.Println("Cold Start...")
	app.Run()
}
