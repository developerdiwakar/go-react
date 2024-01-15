package tokens

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// payload contains the payload (including JWT claims) data of the token
type Payload struct {
	jwt.RegisteredClaims
}

// NewPayload creates a new token payload with specific userid an dcuration
func NewPayload(userId int, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	now := time.Now()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		jwt.RegisteredClaims{
			Audience:  []string{"goreact app"},               // aud
			ID:        tokenId.String(),                      // jti
			IssuedAt:  jwt.NewNumericDate(now),               // iat
			NotBefore: jwt.NewNumericDate(now),               // nbf
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)), // exp
			Subject:   strconv.Itoa(userId),                  // sub
			Issuer:    "Token Authentication",
		},
	}
	return payload, nil
}

// Valid method checks whether the token is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
