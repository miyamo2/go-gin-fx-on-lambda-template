package runner

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/fxapp/runner"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/lambda/handler"
	"go.uber.org/fx"
)

type LambdaRunner struct {
	handler *handler.Handler
}

func (r *LambdaRunner) Run(lifecycle fx.Lifecycle) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting Lambda...")
				defer fmt.Println("Lambda Started")
				go lambda.Start(r.handler.Execute)
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping Lambda")
				return nil
			},
		},
	)
}

func NewLambdaRunner(handler *handler.Handler) runner.FxAppRunner {
	return &LambdaRunner{handler}
}
