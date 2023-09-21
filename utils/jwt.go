package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/alcb1310/bca-go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gitlab.com/0x4149/logz"
)

var ErrExpiredToken = errors.New("token has expired")

const minSecretKeySize = 8

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type JWTMaker struct {
	secretKey string
}

type Maker interface {
	CreateToken(userInfo models.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

func NewJWTMaker(secretKey string) (Maker, error) {
	logz.Debug(len(secretKey), secretKey, minSecretKeySize)
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func NewPayload(u models.User, duration time.Duration) *Payload {
	payload := &Payload{
		ID:        u.Id,
		Email:     u.Email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload
}

func (maker *JWTMaker) CreateToken(userInfo models.User, duration time.Duration) (string, error) {
	payload := NewPayload(userInfo, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	return nil, fmt.Errorf("Not implemented")
}

// func GenerateJWT(u *models.User) (string, error) {
// token := jwt.New(jwt.SigningMethodEdDSA)
// claims := token.Claims.(jwt.MapClaims)
// claims["exp"] = time.Now().Add(10 * time.Minute)
// claims["authorized"] = true
// claims["user"] = u.Email
// tokenString, err := token.SignedString(sampleSecretKey)
// if err != nil {
// return "Signing Error", err
// }
//
// return tokenString, nil
// }
