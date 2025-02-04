package tokenprovider

import (
	"awesomeProject/common"
	"errors"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (Token, error)
	Validate(token string) (TokenPayload, error)
	SecretKey() string
}

type TokenPayload interface {
	UserId() int
	Role() string
}

type Token interface {
	GetToken() string
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ERROR_NOT_FOUND",
		"ERR_NOT_FOUND",
	)
	ErrEncodeToken = common.NewCustomError(
		errors.New("error encode token"),
		"error encode token",
		"ERROR_ENCODE",
		"ERR_ENCODE_TOKEN",
	)

	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token"),
		"invalid token",
		"ERROR_INVALID",
		"ERR_INVALID_TOKEN",
	)
)
