package tokenprovider

import (
	"errors"
	"thucidol/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"), "token not found", "ErrNotFound",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("Error encoding token"), "Error encoding token", "ErrEncodingToken",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("Error invalid token"), "Error invalid token", "ErrInvalidToken",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
