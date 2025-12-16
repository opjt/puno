package api

import (
	"ohp/internal/api/push"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

func NewRouter(pushHandler *push.PushHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/push", pushHandler.Routes())

	return r
}

var routeModule = fx.Module("router",
	fx.Provide(
		push.NewPushHandler,
	),

	fx.Provide(NewRouter),
)
