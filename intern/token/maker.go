package token

import "time"

type Maker interface {
	CreateToken(username string, userId int64, role string, duration time.Duration) (string, Payload, error)
	VerifyToken(token string) (*Payload, error)
}
