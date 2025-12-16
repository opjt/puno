package push

import (
	"net/http"
	"ohp/internal/pkg/log"

	"github.com/go-chi/chi/v5"
)

type PushHandler struct {
	log *log.Logger
}

func NewPushHandler(log *log.Logger) *PushHandler {
	return &PushHandler{
		log: log,
	}
}
func (h *PushHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", h.Ping)
	return r
}

func (h *PushHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("pong"))
	if err != nil {
		h.log.Warn("Failed to write response", "err", err)
	}
}
