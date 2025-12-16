package v1

import (
	"ohp/internal/server/httpr/v1/handler"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type RouteParams struct {
	fx.In
	PushHandler *handler.PushHandler
}

func newRouter(params RouteParams) chi.Router {
	r := chi.NewRouter()

	r.Route("/push", func(r chi.Router) {
		r.Get("/", params.PushHandler.Ping)
	})

	return r
}

var Module = fx.Module("v1-api",
	fx.Provide(handler.NewPushHandler),

	fx.Provide(
		fx.Annotate(
			newRouter,
			fx.ResultTags(`name:"v1"`),
		),
	),
)
