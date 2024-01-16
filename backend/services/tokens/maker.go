package tokens

import "time"

// Maker is an interface to manage token
type Maker interface {
	// CreateToken create a new token for a specific userid and duration
	CreateToken(userId int, duration time.Duration) (string, error)
	// VerifyToken verify if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
