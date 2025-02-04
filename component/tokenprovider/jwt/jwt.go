package jwt

import (
	"awesomeProject/common"
	"awesomeProject/component/tokenprovider"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtProvider struct {
	secret string
	prefix string
}

func NewTokenJWTProvider(prefix string) *jwtProvider {
	return &jwtProvider{prefix: prefix}
}

type myClaims struct {
	Payload common.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

type token struct {
	Token   string    `json:"token"`
	Expiry  int       `json:"expiry"`
	Created time.Time `json:"created"`
}

func (t *token) GetToken() string {
	return t.Token
}

func (j *jwtProvider) Generate(data common.TokenPayload, expiry int) (tokenprovider.Token, error) {
	now := time.Now()

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		myClaims{
			common.TokenPayload{
				UId:   data.GetUId(),
				URole: data.GetURole(),
			},
			jwt.StandardClaims{
				ExpiresAt: now.Local().Add(time.Second * time.Duration(expiry)).Unix(),
				IssuedAt:  now.Local().Unix(),
				Id:        fmt.Sprintf("%d", now.UnixNano()),
			},
		})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}
	return &token{Token: myToken, Expiry: expiry, Created: now}, nil
}
