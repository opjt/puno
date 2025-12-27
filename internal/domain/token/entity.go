package token

import "github.com/google/uuid"

type Token struct {
	ID       uuid.UUID
	P256dh   string
	Auth     string
	UserID   uuid.UUID
	EndPoint string
}
