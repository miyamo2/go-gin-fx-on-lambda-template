package gin

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/fxapp/runner"
	"go.uber.org/fx"
)

type GinRunner struct {
	engine *gin.Engine
	uri    []string
}

func (r *GinRunner) Run(lifecycle fx.Lifecycle) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				fmt.Println("Starting Gin")
				if len(r.uri) == 1 {
					go r.engine.Run(r.uri[0])
				} else {
					go r.engine.Run()
				}
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping Gin")
				return nil
			},
		},
	)
}

func NewGinRunner(engine *gin.Engine, uri ...string) runner.FxAppRunner {
	uri_len := len(uri)
	if uri_len > 1 {
		panic(fmt.Sprintf("Too many arguments for NewGinRunner: %d", uri_len))
	}
	return &GinRunner{engine, uri}
}
