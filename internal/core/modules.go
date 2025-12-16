package core

import (
	"ohp/internal/api"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	api.Module,
)
