package endpoint

import (
	"context"
	"ohp/internal/pkg/token"
)

type EndpointService struct {
	repo EndpointRepository
}

func NewEndpointService(
	repo EndpointRepository,
) *EndpointService {
	return &EndpointService{
		repo: repo,
	}
}

func (s *EndpointService) Add(ctx context.Context, serviceName string) error {

	userClaim, err := token.UserFromContext(ctx)
	if err != nil {
		return err
	}

	if err := s.repo.Add(ctx, insertEndpointParams{
		userID:      userClaim.UserID,
		serviceName: serviceName,
		endpoint:    "test",
	}); err != nil {
		return err
	}

	return nil
}

// func GenerateEndpointToken(length int) (string, error) {

// 	byteLen := length * 6 / 8
// 	if byteLen < 8 {
// 		byteLen = 8
// 	}

// 	b := make([]byte, byteLen)
// 	if _, err := rand.Read(b); err != nil {
// 		return "", err
// 	}

// 	token := base62Encode(b)
// 	if len(token) > length {
// 		token = token[:length]
// 	}
// 	for len(token) < length {
// 		token = "0" + token
// 	}
// 	return token, nil
// }

// const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// func base62Encode(b []byte) string {
// 	var result []byte
// 	var num = new(big.Int).SetBytes(b)
// 	base := big.NewInt(62)
// 	zero := big.NewInt(0)
// 	mod := new(big.Int)

// 	for num.Cmp(zero) > 0 {
// 		num.DivMod(num, base, mod)
// 		result = append(result, base62Chars[mod.Int64()])
// 	}
// 	// reverse
// 	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
// 		result[i], result[j] = result[j], result[i]
// 	}

// 	return string(result)
// }
