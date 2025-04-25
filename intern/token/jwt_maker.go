package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(key string) (*PasetoMaker, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key size: must be exactly 32 bytes")
	}

	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(key),
	}, nil
}

func (maker *PasetoMaker) CreateToken(username string, userId int64, role string, duration time.Duration) (string, Payload, error) {

	payload, err := NewPayload(username, userId, role, duration)
	if err != nil {
		return "", *payload, fmt.Errorf("failed to create payload: %w", err)
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return "", *payload, fmt.Errorf("failed to encrypt token: %w", err)
	}

	return token, *payload, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	var payload Payload
	err := maker.paseto.Decrypt(token, maker.symmetricKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
