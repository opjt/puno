package push

import (
	"context"
	"encoding/json"
	"fmt"
	"ohp/internal/domain/endpoint"
	"ohp/internal/domain/token"
	"ohp/internal/pkg/config"

	"github.com/SherClockHolmes/webpush-go"
)

type PushService struct {
	// repo     SubscriptionRepository
	vapidKey config.Vapid

	tokenService    *token.TokenService
	endpointService *endpoint.EndpointService
}

func NewPushService(
	// repo SubscriptionRepository,
	env config.Env,

	tokenService *token.TokenService,
	endpointService *endpoint.EndpointService,
) *PushService {
	return &PushService{
		// repo:            repo,
		vapidKey:        env.Vapid,
		tokenService:    tokenService,
		endpointService: endpointService,
	}
}

func (s *PushService) Subscribe(ctx context.Context, sub Subscription) error {

	if err := s.tokenService.Register(ctx, token.Token{
		P256dh:   sub.P256dh,
		Auth:     sub.Auth,
		UserID:   sub.UserID,
		EndPoint: sub.Endpoint,
	}); err != nil {
		return err
	}

	return nil
}

func (s *PushService) Unsubscribe(ctx context.Context, sub Subscription) error {

	if err := s.tokenService.Unregister(ctx, token.Token{
		P256dh:   sub.P256dh,
		Auth:     sub.Auth,
		EndPoint: sub.Endpoint,
	}); err != nil {
		return err
	}

	return nil
}

// Push notification using endpoint token
func (s *PushService) Push(ctx context.Context, endpointToken string) error {

	endpoint, err := s.endpointService.FindByToken(ctx, endpointToken)
	if err != nil {
		return err
	}
	userID := endpoint.UserID

	tokens, err := s.tokenService.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		if err := s.pushNotification(token, "title1", "body4"); err != nil {
			return err
		}
	}

	return nil
}

func (s *PushService) pushNotification(token token.Token, title, body string) error {

	subs := &webpush.Subscription{
		Endpoint: token.EndPoint,
		Keys: webpush.Keys{
			P256dh: token.P256dh,
			Auth:   token.Auth,
		},
	}

	options := &webpush.Options{
		VAPIDPublicKey:  s.vapidKey.PublicKey,
		VAPIDPrivateKey: s.vapidKey.PrivateKey,
		TTL:             300,
		Subscriber:      "jtpark1957@gmail.com",
	}
	payload := map[string]interface{}{
		"title": title,
		"body":  body,
		"data": map[string]string{
			"url":       "/",
			"timestamp": fmt.Sprintf("%d", 1234567890),
		},
	}

	payloadBytes, _ := json.Marshal(payload)

	resp, err := webpush.SendNotification(payloadBytes, subs, options)
	if err != nil {
		return err
	}
	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil

}
