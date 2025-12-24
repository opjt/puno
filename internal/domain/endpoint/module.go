package endpoint

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewEndpointService),
	fx.Provide(NewEndpointRepository),
)
