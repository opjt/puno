package handler

import (
	"context"
	"ohp/internal/api/wrapper"
	"ohp/internal/domain/endpoint"
	"ohp/internal/pkg/log"

	"github.com/go-chi/chi/v5"
)

type EndpointHandler struct {
	log     *log.Logger
	service *endpoint.EndpointService
}

func NewEndpointHandler(log *log.Logger, service *endpoint.EndpointService) *EndpointHandler {
	return &EndpointHandler{
		log:     log,
		service: service,
	}
}
func (h *EndpointHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", wrapper.WrapJson(h.Add, h.log.Error, wrapper.RespondJSON))
	return r
}

type reqAddEndpoint struct {
	ServiceName string `json:"serviceName"`
}

func (h *EndpointHandler) Add(ctx context.Context, req reqAddEndpoint) (interface{}, error) {

	h.log.Info("req", "sub", req)
	if err := h.service.Add(ctx, req.ServiceName); err != nil {
		return nil, err
	}

	return "success", nil
}
