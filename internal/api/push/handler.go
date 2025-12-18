package push

import (
	"context"
	"ohp/internal/api/wrapper"
	"ohp/internal/domain/push"
	"ohp/internal/pkg/log"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/go-chi/chi/v5"
)

type PushHandler struct {
	log *log.Logger

	repo push.SubscriptionRepository // fx가 자동으로 주입해줌
}

func NewPushHandler(log *log.Logger, repo push.SubscriptionRepository) *PushHandler {
	return &PushHandler{
		log:  log,
		repo: repo,
	}
}
func (h *PushHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", wrapper.WrapJson(h.Subscribe, h.log.Error, wrapper.RespondJSON))
	return r
}

type subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

// Subscribe, push subscribe
func (h *PushHandler) Subscribe(ctx context.Context, req subscription) (interface{}, error) {

	subs := &webpush.Subscription{
		Endpoint: req.Endpoint,
		Keys: webpush.Keys{
			P256dh: req.Keys.P256dh,
			Auth:   req.Keys.Auth,
		},
	}
	h.repo.Save(req.Endpoint, subs)
	h.log.Info("push subscribe", "endpoint", req.Endpoint)
	return "success1", nil
}

// Push, notification
func (h *PushHandler) Push(ctx context.Context, req any) error {
	return nil
}
func (h *PushHandler) Broadcast(ctx context.Context, req any) error {
	return nil
}
