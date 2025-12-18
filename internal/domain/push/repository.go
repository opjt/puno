package push

import (
	"sync"

	"github.com/SherClockHolmes/webpush-go"
)

// 추후 DB 구현체로 교체 가능하도록 인터페이스 정의
type SubscriptionRepository interface {
	Save(id string, sub *webpush.Subscription)
	Get(id string) (*webpush.Subscription, bool)
}

type MemorySubscriptionRepo struct {
	mu            sync.RWMutex
	subscriptions map[string]*webpush.Subscription
}

func NewMemorySubscriptionRepo() SubscriptionRepository {
	return &MemorySubscriptionRepo{
		subscriptions: make(map[string]*webpush.Subscription),
	}
}

func (m *MemorySubscriptionRepo) Save(id string, sub *webpush.Subscription) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.subscriptions[id] = sub
}

func (m *MemorySubscriptionRepo) Get(id string) (*webpush.Subscription, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	sub, ok := m.subscriptions[id]
	return sub, ok
}
