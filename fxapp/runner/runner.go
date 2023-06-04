package runner

import "go.uber.org/fx"

type FxAppRunner interface {
	Run(lifecycle fx.Lifecycle)
}
