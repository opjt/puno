package pkg

import (
	"ohp/internal/pkg/config"
	"ohp/internal/pkg/token"
	"time"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		func(env config.Env) *token.TokenProvider {
			// 직접 아규먼트 값 주입 TODO: env화
			return token.NewTokenProvider(
				env.JWTSecret, // secret
				"ohp-api",     // issuer
				24*time.Hour,  // expiry
			)
		},
	),
)
