package handler

import (
	"net/http"
)

type PushHandler struct{}

func NewPushHandler() *PushHandler {
	return &PushHandler{}
}

func (h *PushHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
