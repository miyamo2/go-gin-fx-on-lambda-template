package routing

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(createGinEngineFromHelloController))
