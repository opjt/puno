package httpr

import (
	v1 "ohp/internal/server/httpr/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

type RouterParams struct {
	fx.In
	V1 chi.Router `name:"v1"`
}

func NewRouter(params RouterParams) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/api/v1", params.V1)

	return r
}

var routeModule = fx.Module("router",
	v1.Module,
	fx.Provide(NewRouter),
)
