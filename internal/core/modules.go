package core

import (
	"ohp/internal/server/httpr"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	httpr.Module,
)
