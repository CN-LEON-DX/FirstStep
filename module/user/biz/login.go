package biz

import (
	"awesomeProject/common"
	"awesomeProject/component/tokenprovider"
	"awesomeProject/module/user/model"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfor ...string) (*model.User, error)
}

type loginBussiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBussiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginBussiness {
	return &loginBussiness{storeUser: storeUser, tokenProvider: tokenProvider, hasher: hasher, expiry: expiry}
}

// 1. Find user, email
// 2. Hash password from input and compare with password in db
// 3. Provider issue JWT token for client
// 3.1 Access token and refresh token
// 4. Return token to client

func (bussinesss *loginBussiness) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	user, err := bussinesss.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	passHashed := bussinesss.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	accessToken, err := bussinesss.tokenProvider.Generate(payload, bussinesss.expiry)

	if err != nil {
		return nil, common.NewCustomError(err, "cannot generate token", "loginBussiness", "Generate")
	}

	// refreshToken, err := bussinesss.tokenProvider.Generate(payload, bussinesss.tkCfg.GetRtExp())

	return accessToken, nil
}
