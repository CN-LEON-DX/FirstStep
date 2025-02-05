package biz

import (
	"awesomeProject/common"
	"awesomeProject/module/user/model"
	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}, moreInfor ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(password string) string
}

type registerBussiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBussiness(registerStorage RegisterStorage, hasher Hasher) *registerBussiness {
	return &registerBussiness{registerStorage: registerStorage, hasher: hasher}
}

func (bussiness *registerBussiness) Register(ctx context.Context, data *model.UserCreate) error {
	user, _ := bussiness.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		//if user.Status == 0 {
		//	return user has been disable
		//}

		return model.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = bussiness.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	if err := bussiness.registerStorage.CreateUser(ctx, data); err != nil {
		return common.CannotGetEntity(model.User{}.TableName(), err)
	}
	return nil
}
